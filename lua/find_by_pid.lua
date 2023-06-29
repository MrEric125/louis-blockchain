-- 检查表中是否存在指定元素
function table.contains(table, element)
    for _, value in pairs(table) do
        if value == element then
            return true
        end
    end
    return false
end

-- 获取部门ids
data = Array(data)
nodeIds = data:arrayColumn('id')

-- 部门id与用户id对应关系
function getNodeId2UserIds(nodeIds)
    local result = {}
    if nodeIds == nil or nodeIds == {} or table.getn(nodeIds) == 0 then
        return result
    end
    local sql = [[
	    select node_id, id from user where is_deleted = 0 and biz_id = {bizId} and node_id in ({nodeIds});
    ]]

    local params = {
        bizId = user.bizId,
        nodeIds = string.join(",", nodeIds)
    }
    local fetch_res = mysql_fetch_all('scrm', sql % params)
    if fetch_res then
        for k,v in pairs(fetch_res) do
            local nodeId = v.node_id
            local userId = v.id
            if not result[nodeId] then
                result[nodeId] = {}
            end
            table.insert(result[nodeId],userId)
        end
    end
    return result
end

    

-- 获取某个部门下的一批成员对应的绑定线索数量
function selectBindCount(userIds)
    local result = 0
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
        SELECT
    	    count( 0 ) count
        FROM
    	    {crm_schema_name}.leads_binding_info
        WHERE
        	biz_id = {bizId} 
        	and binding_time BETWEEN {startTime} AND {endTime}
        	AND binding_user_id IN ({userIds}) 
    ]]

    local params = {
        crm_schema_name = crm_schema_name,
        bizId = user.bizId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userIds = string.join(",", userIds)
    }
    local fetch_res = mysql_fetch_all('crm', sql % params)
    if fetch_res then
       return fetch_res[0].count
    end
    return result
end

-- 获取用户对应的手机号
function selectNewFollowMobiles(userIds)
    local result = {}
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
    SELECT
	GROUP_CONCAT( ll.mobile ) mobiles,
	binding_user_id userId 
FROM
	{crm_schema_name}.leads_binding_info lbi
	LEFT JOIN {crm_schema_name}.leads_lib ll ON lbi.leads_id = ll.id 
WHERE
	lbi.biz_id = {bizId} 
	and binding_time BETWEEN {startTime} AND {endTime}
	AND binding_user_id IN ({userIds}) 
GROUP BY
	binding_user_id]]
    
    local params = {
        crm_schema_name = crm_schema_name,
        bizId = user.bizId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userIds = string.join(",", userIds)
    }
    
    local fetch_res = mysql_fetch_all('crm', sql % params)
    if fetch_res then
        for k,v in pairs(fetch_res) do
            result[v.userId] = v.mobiles
        end
    end
    return result;
end

-- 获取用户对应的新跟进量
function selectNewFollowCount(userId, mobiles)
    local result = {}
    if mobiles == nil then 
        result["count"] = 0;
        return result
    end
    local sql = [[
SELECT customerPhone
FROM
	{index}
WHERE
	bizId = {bizId}
	AND userId = {userId}
	AND customerPhone IN ({mobiles})
	and lastTimeStamp >= {time1}
	and lastTimeStamp <= {time2}
	]]
    
    local year1 = string.sub(getStartTime(), 0, 4)
    local month1 = string.sub(getStartTime(), 6, 7)
    local day1 = string.sub(getStartTime(), 9, 10)
    local year2 = string.sub(request.endTime, 0, 4)
    local month2 = string.sub(request.endTime, 6, 7)
    local day2 = string.sub(request.endTime, 9, 10)
    local toTime1 = os.time({year = year1, month = month1, day = day1, hour = 00, min = 00 , sec = 00})
    local toTime2 = os.time({year = year2, month = month2, day = day2, hour = 23, min = 59 , sec = 59})
    local params = {
        index = follow_index,
        bizId = user.bizId,
        time1 = toTime1 * 1000,
        time2 = toTime2 * 1000,
        userId = userId,
        mobiles = "'" .. mobiles:gsub(",", "','") .. "'"
    }
    local fetch_res = es_fetch_all('elastic', sql % params)
    local customerPhones = {}
    local es_count = 0
    for _, record in pairs(fetch_res) do
        local customerPhone = record.customerPhone
        if customerPhone ~= nil and not table.contains(customerPhones, customerPhone) then
            table.insert(customerPhones, customerPhone)
            es_count = es_count + 1
        end
    end
    -- 从 mobiles 中去掉 customerPhones
    local mobiles_table = {}
    for _, mobile in ipairs(string.split(mobiles, ",")) do
        mobile = string.trim(mobile)
        if (customerPhones == nil or next(customerPhones) == nil) or (mobile ~= "" and not table.contains(customerPhones, mobile)) then
            table.insert(mobiles_table, mobile)
        end
    end
    local crm_sql = [[
        select count(distinct ll.id) count  from {schema}.leads_binding_info lbi
        join {schema}.leads_follow_record lfr on lbi.leads_id = lfr.leads_id  and lfr.create_time between {startTime} and {endTime}
        join {schema}.leads_lib ll on lbi.leads_id = ll.id
        where lbi.binding_user_id = {userId} and ll.mobile in ({mobiles})
    ]]
    local crm_params = {
        schema = crm_schema_name,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userId = userId,
        mobiles = "'" .. table.concat(mobiles_table, "','") .. "'"
    }
    local crm_fetch_result = mysql_fetch_all('crm', crm_sql % crm_params);
    result["count"] = es_count + crm_fetch_result[0].count
    return result
end

-- 获取用户id与weworkUserId对应关系
function selectUserIdWeworkUserId(userIds)
    local result = {}
    local sql = [[
        select id, wework_user_id from user where id in ({userIds})
	]]
	
	local params = {
	    userIds = string.join(",", userIds)
	}
	local fetch_res = mysql_fetch_all('scrm', sql % params)
	for _, row in pairs(fetch_res) do
	    result[row.id] = row.wework_user_id
	end
	return result
end

-- 获取加v数，聊天数
function getAddWeworkCount(userId, weworkUserId)
    local result = {}
    local res = {}
    res.addWeworkCount = 0
    res.weworkChatCount = 0
    if userId == nil or weworkUserId == nil then
        return res
    end
    -- 获取有contactId的线索
    local sql = [[
        select cl.contact_id from {crm_schema_name}.leads_binding_info lbi join {crm_schema_name}.leads_lib ll on ll.id = lbi.leads_id
        join {crm_schema_name}.customer_lib cl on cl.id = ll.customer_id where lbi.binding_user_id = {userId} and cl.contact_id != ''
	    and lbi.binding_time BETWEEN {startTime} AND {endTime}
	]]
    

    local params = {
        crm_schema_name = crm_schema_name,
        userId = userId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
    }
    local fetch_res = mysql_fetch_all('crm', sql % params)
    if fetch_res then
        for k,v in pairs(fetch_res) do
            table.insert(result, v.contact_id)
        end
    end
    
    -- 获取是好友关系的数量
    local conversionIds = {}
    if fetch_res ~= nil and next(fetch_res) ~= nil then
        local scrm_sql = [[
            select contact_id from wework_contact_relation where wework_user_id = '{weworkUserId}' and contact_id in ({contactIds}) and is_deleted = 0
        ]]
        local scrm_params = {
            weworkUserId = weworkUserId,
            contactIds = "'" .. table.concat(result, "','") .. "'"
        }
        local scrm_fetch_res = mysql_fetch_all('scrm', scrm_sql % scrm_params)
        
        if scrm_fetch_res then
            for k,v in pairs(scrm_fetch_res) do
                res.addWeworkCount = res.addWeworkCount + 1
                table.insert(conversionIds, weworkUserId .."$$"..v.contact_id)
                table.insert(conversionIds, v.contact_id .."$$"..weworkUserId)
            end
        end
    end
    -- 获取有聊天的数量
    if conversionIds ~= nil and next(conversionIds) ~= nil then
        local es_sql = [[
select count(distinct(conversationId.keyword)) as count from {index} where conversationId.keyword in ({conversionIds}) and msgTime between {time1} and {time2}]]

        local year1 = string.sub(getStartTime(), 0, 4)
        local month1 = string.sub(getStartTime(), 6, 7)
        local day1 = string.sub(getStartTime(), 9, 10)
        local year2 = string.sub(request.endTime, 0, 4)
        local month2 = string.sub(request.endTime, 6, 7)
        local day2 = string.sub(request.endTime, 9, 10)
        local toTime1 = os.time({year = year1, month = month1, day = day1, hour = 00, min = 00 , sec = 00})
        local toTime2 = os.time({year = year2, month = month2, day = day2, hour = 23, min = 59 , sec = 59})
        local es_params = {
            index = chat_message_index,
            time1 = toTime1 * 1000,
            time2 = toTime2 * 1000,
            conversionIds = "'" .. table.concat(conversionIds, "','") .. "'"
        }
        local es_fetch_res = es_fetch_all('elastic', es_sql % es_params)
        if es_fetch_res then
           res.weworkChatCount = es_fetch_res[0].count
        end    
    end   
    return res
end

-- 总跟进客户数
function selectFollowCustomerCount(userIds)
    local result = 0
    local userId2Mobiles = {}
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
SELECT distinct
	customerPhone,
    userId 
FROM
	{index}
WHERE
	bizId = {bizId} 
	AND userId IN ({userIds})
	and lastTimeStamp >= {time1}
	and lastTimeStamp <= {time2}]]
    

    local year1 = string.sub(getStartTime(), 0, 4)
    local month1 = string.sub(getStartTime(), 6, 7)
    local day1 = string.sub(getStartTime(), 9, 10)
    local year2 = string.sub(request.endTime, 0, 4)
    local month2 = string.sub(request.endTime, 6, 7)
    local day2 = string.sub(request.endTime, 9, 10)
    local toTime1 = os.time({year = year1, month = month1, day = day1, hour = 00, min = 00 , sec = 00})
    local toTime2 = os.time({year = year2, month = month2, day = day2, hour = 23, min = 59 , sec = 59})
    local params = {
        index = follow_index,
        bizId = user.bizId,
        time1 = toTime1 * 1000,
        time2 = toTime2 * 1000,
        userIds = string.join(",", userIds)
    }

    local fetch_res = es_fetch_all('elastic', sql % params)
    if fetch_res then
        for k,v in pairs(fetch_res) do
            if userId2Mobiles[v.userId] == nil then
                userId2Mobiles[v.userId] = {}
            end
            if not table.contains(userId2Mobiles[v.userId], v.customerPhone) then
                result = result + 1
                table.insert(userId2Mobiles[v.userId], v.customerPhone)
            end
        end
    end
    
    -- 查询leads_follow_record
    local crm_params = {
        schema = crm_schema_name,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
    }
    
    for k, v in pairs(userIds) do 
        local crm_sql = [[
            select count(distinct lfr.leads_id) count  from {schema}.leads_binding_info lbi
            join {schema}.leads_follow_record lfr on lbi.leads_id = lfr.leads_id  and lfr.create_time between {startTime} and {endTime}
            {joinLeadsLib} where lbi.binding_user_id = {userId} {notIncludeMobiles}
        ]]

        local userIdStr = tostring(v)
        crm_params["userId"] = v
        if userId2Mobiles[userIdStr] then 
            local joinLeadsLib = [[join {schema}.leads_lib ll on lbi.leads_id = ll.id]]
            
            local notIncludeMobiles = [[and ll.mobile not in ({mobiles})]]
            crm_sql = crm_sql:gsub("{joinLeadsLib}", joinLeadsLib)
            crm_sql = crm_sql:gsub("{notIncludeMobiles}", notIncludeMobiles)
            crm_params["mobiles"] = "'" .. table.concat(userId2Mobiles[userIdStr], "','") .. "'"
        else
            crm_sql = crm_sql:gsub("{joinLeadsLib}", "")
            crm_sql = crm_sql:gsub("{notIncludeMobiles}", "")
        end
        local crm_fetch_result = mysql_fetch_all('crm', crm_sql % crm_params);
       
        if crm_fetch_result then
           result = result + crm_fetch_result[0].count
        end         
    end
    return result
end

-- 总跟进次数
function selectFollowCount(userIds)
    local result = 0
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
SELECT
	count( * ) count
FROM
	{index}
WHERE
	bizId = {bizId} 
	AND userId IN ({userIds}) 
	and lastTimeStamp >= {time1}
	and lastTimeStamp <= {time2}
]]
    

    local year1 = string.sub(getStartTime(), 0, 4)
    local month1 = string.sub(getStartTime(), 6, 7)
    local day1 = string.sub(getStartTime(), 9, 10)
    local year2 = string.sub(request.endTime, 0, 4)
    local month2 = string.sub(request.endTime, 6, 7)
    local day2 = string.sub(request.endTime, 9, 10)
    local toTime1 = os.time({year = year1, month = month1, day = day1, hour = 00, min = 00 , sec = 00})
    local toTime2 = os.time({year = year2, month = month2, day = day2, hour = 23, min = 59 , sec = 59})
    local params = {
        index = follow_index,
        bizId = user.bizId,
        time1 = toTime1 * 1000,
        time2 = toTime2 * 1000,
        userIds = string.join(",", userIds)
    }
    
    local fetch_res = es_fetch_all('elastic', sql % params)
    if fetch_res then
        result = fetch_res[0].count
    end
    
    -- 查询leads_follow_record表
    local crm_sql = [[
        select count(1) count
        from {schema}.leads_binding_info lbi join {schema}.leads_follow_record lfr on lbi.leads_id = lfr.leads_id
        and lbi.binding_user_id in ({userIds}) and lfr.create_time between {startTime} and {endTime}
        group by lbi.binding_user_id;
    ]]
    local crm_params = {
        schema = crm_schema_name,
        bizId = user.bizId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userIds = string.join(",", userIds)
    }
    local crm_fetch_result = mysql_fetch_all('crm', crm_sql % crm_params);
    if crm_fetch_result ~= nil and next(crm_fetch_result) ~= nil then
        result = result + crm_fetch_result[0].count
    end
    return result
end

-- 外呼次数
function selectCallCount(userIds)
    local result = {}
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
	SELECT
    user_id userId,
	count( 0 ) count,
	IFNULL( sum( STATUS = 1 ), 0 ) successCount,
	sum( call_duration ) callDuration
FROM
	call_record
WHERE biz_id = {bizId} 
	and create_time BETWEEN {startTime} AND {endTime}
	AND user_id IN ({userIds})
GROUP BY
	user_id;
    ]]

    local params = {
        bizId = user.bizId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userIds = string.join(",", userIds)
    }
    local fetch_res = mysql_fetch_all('scrm', sql % params)
    return fetch_res;

end

-- 发送短信数量
function selectSmsCount(userIds)
    local result = 0
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
	SELECT
	count( 0 ) count
FROM
	sms_task_detail 
WHERE
	biz_id = {bizId} 
	and create_time BETWEEN {startTime} AND {endTime}
	AND create_by IN ({userIds})
    ]]

    local params = {
        bizId = user.bizId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userIds = string.join(",", userIds)
    }
    local fetch_res = mysql_fetch_all('scrm', sql % params)
    if fetch_res then
        result = fetch_res[0].count
    end
    return result
end

-- 获取订单信息
function selectOrderCount(userIds)
    local result = {}
    if userIds == nil or userIds == {} or table.getn(userIds) == 0 then
        return result
    end
    local sql = [[
	SELECT
	count( 0 ) count,
	count( distinct customer_num  ) paidCount,
	IFNULL( sum( paid_amount ), 0 ) paidAmount
FROM
	{schema}.customer_order 
WHERE
	biz_id = {bizId} 
	and create_time BETWEEN {startTime} AND {endTime}
	AND user_id IN ({userIds})
	and pay_status = 3
    ]]

    local params = {
        schema = '`' .. customer_schema .. '`',
        bizId = user.bizId,
        startTime = "'" .. getStartTime() .. " 00:00:00'",
        endTime = "'" .. request.endTime .. " 23:59:59'",
        userIds = string.join(",", userIds)
    }
    local fetch_res = mysql_fetch_all('customer', sql % params)
    return fetch_res;
end

local nodeId2UserIds = getNodeId2UserIds(nodeIds)



local nodeTreeMap=getNodeTreeMap(nodeIds)

-- 获取部门下的子部门




for k, v in pairs(data) do
    v.name = v["name"]
    local nodeId = v["id"]
    local progFormat = "%.2f"
    
    v.userBindCount = 0
    v.userNewFollowCount = 0
    v.addWeworkCount = 0
    v.weworkChatCount = 0
    v.userFollowCustomerCount = 0
    v.userFollowCount = 0
    v.callCount = 0
    v.callDuration = 0
    v.callRate = 0
    v.smsCount = 0
    v.paidCount = 0
    v.orderCount = 0
    v.paidAmount = 0
    
    -- 新分配客户数
    if nodeId2UserIds ~= nil then
        userIds = nodeId2UserIds[nodeId]
        if userIds ~= nil then
            -- 新分配客户
            v.userBindCount = selectBindCount(userIds)
            local userId2Mobiles = selectNewFollowMobiles(userIds)   
            
            local userNewFollowCount = 0
            local addWeworkCount = 0
            local weworkChatCount = 0
            if userId2Mobiles ~= nil then
                for k, v in pairs(userId2Mobiles) do 
                    local userIdNewFollowCount = selectNewFollowCount(k, v)
                    -- 跟进新客户数
                    if userIdNewFollowCount ~= nil then
                        userNewFollowCount = userNewFollowCount + userIdNewFollowCount["count"]
                    end
                end
                v.userNewFollowCount = userNewFollowCount
            end
            
            userId2WeworkUserId = selectUserIdWeworkUserId(userIds)
            for k, v in pairs(userIds) do 
                if userId2WeworkUserId ~= nil and userId2WeworkUserId ~= {} and userId2WeworkUserId[v] ~= nil then
                    local addChatRes = getAddWeworkCount(v, userId2WeworkUserId[v])
                    addWeworkCount = addWeworkCount +  addChatRes.addWeworkCount
                    weworkChatCount = weworkChatCount + math.floor(addChatRes.weworkChatCount)
                end    
            end
            v.addWeworkCount = addWeworkCount
            v.weworkChatCount = weworkChatCount
            
            v.userFollowCustomerCount = selectFollowCustomerCount(userIds)
            v.userFollowCount = math.floor(selectFollowCount(userIds))
            
            userCallCount = selectCallCount(userIds)
            local callCount = 0
            local totalDuration = 0
            local callSuccessCount = 0
            if userCallCount ~= nil then
                for j,k in pairs(userCallCount) do
                    callCount = callCount + k.count
                    totalDuration = totalDuration + k.callDuration
                    callSuccessCount = callSuccessCount + k.successCount
                    
                end
            end
            v.callCount = callCount
            local minutes = math.floor(totalDuration / 60)
            local seconds = math.mod(totalDuration, 60)
            if seconds < 0 then
                seconds = 0
            end
            
            if minutes > 0 or seconds > 0 then
                v.callDuration = minutes .. '分' .. seconds .. '秒'
            end
            v.callSuccessCount=callSuccessCount
            if callCount > 0 then
                v.callRate = callSuccessCount / callCount
                v.callRate = progFormat:format(v.callRate * 100.0) .. "%"
            end
        end 
        
        v.smsCount = selectSmsCount(userIds)
        
        local userOrderCount = selectOrderCount(userIds)
        if userOrderCount ~= nil and userOrderCount[0] ~= nil then 
            v.paidCount = userOrderCount[0].paidCount
            v.orderCount = userOrderCount[0].count
            v.paidAmount = userOrderCount[0].paidAmount / 10000
        end
        
    end
    
    
    if v.userBindCount == 0 then
        v.followRate = "--"
        v.addWeworkRate = "--"
    else
        v.followRate = progFormat:format(v.userNewFollowCount / v.userBindCount * 100.0) .. "%"
        v.addWeworkRate = progFormat:format(v.addWeworkCount / v.userBindCount * 100.0) .. "%"
    end
end




for i = #data, 1, -1 do
    local v = data[i]
     local progFormat = "%.2f"
    

    
    for _,withChildNodes in pairs(nodeTreeMap) do
       
        
        if withChildNodes ~=nill and withChildNodes~={} then
            --  有 子node 的数据，
          if  withChildNodes.id==v.id and withChildNodes.ids ~=nill and #withChildNodes.ids ~=0 then
             
            
                -- 找到子节点对应数，做数据拼接
                local childIds=withChildNodes.ids
                 for _,childId in pairs(childIds) do

                     for _,childData in pairs(data) do
                          
                          
                          if childData ~=nill and childData ~={}  and childId ~=nil and childData.id==childId then 
                              
                                 v.userFollowCustomerCount=v.userFollowCustomerCount+childData.userFollowCustomerCount
                                 v.callCount=v.callCount+childData.callCount
                                
                                 v.addWeworkCount=v.addWeworkCount+childData.addWeworkCount
                                
                                 v.paidAmount=v.paidAmount+childData.paidAmount
                                 v.userBindCount=v.userBindCount+childData.userBindCount
                                
                                 v.userNewFollowCount=v.userNewFollowCount+childData.userNewFollowCount
                                 v.paidCount=v.paidCount+childData.paidCount
                                
                                 v.weworkChatCount=v.weworkChatCount+childData.weworkChatCount
                                 v.userFollowCount=v.userFollowCount+childData.userFollowCount
                                 v.smsCount=v.smsCount+childData.smsCount
                                 v.orderCount=v.orderCount+childData.orderCount
                             
                          end
                         
                     end
                     
                 end
             
          end
              
        end
        
    end
    
   


    
    if v.userBindCount == 0 then
        v.followRate = "--"
        v.addWeworkRate = "--"
    else
        v.followRate = progFormat:format(v.userNewFollowCount / v.userBindCount * 100.0) .. "%"
        v.addWeworkRate = progFormat:format(v.addWeworkCount / v.userBindCount * 100.0) .. "%"
    end
    
     if callCount~=nil and  callCount > 0 then
                v.callRate = callSuccessCount / callCount
                v.callRate = progFormat:format(v.callRate * 100.0) .. "%"
     end
    
    
    
    if v.userBindCount == 0 and v.userNewFollowCount == 0 and v.addWeworkCount == 0 
        and v.weworkChatCount == 0 and v.userFollowCustomerCount == 0 and v.userFollowCount == 0 
        and v.callCount == 0 and v.smsCount == 0 and v.paidCount == 0 and v.orderCount == 0 
        and v.paidAmount == 0 then
        table.remove(data, i)
        end
    

end

keys = {
   "name",  "userBindCount", "userNewFollowCount", "followRate", "addWeworkCount", "addWeworkRate", "userFollowCustomerCount","userFollowCount", 
   "weworkChatCount", "callCount","callRate","callDuration","smsCount","paidCount","orderCount","paidAmount"
}
return {
    keys = keys,
    data = data,
    schema = {
        name = '部门',
        userBindCount = '新分配客户',
        userNewFollowCount = '跟进新客户数',
        followRate = '跟进率',
        addWeworkCount = '新客户加微数',
        addWeworkRate = '加微率',
        userFollowCustomerCount = '总跟进客户数',
        userFollowCount = '总跟进次数',
        weworkChatCount = '微聊客户数',
        callCount = '外呼次数',
        callRate = '外呼接通率',
        callDuration = '通话时长(分)',
        smsCount = '发短信',
        paidCount = '成交客户',
        orderCount = '订单量',
        paidAmount = '订单支付金额(元)'
    },
    style = {
        width = {
            name = '100px',
            userBindCount = '100px',
            userNewFollowCount = '120px',
            followRate = '100px',
            addWeworkCount = '120px',
            addWeworkRate = '100px',
            userFollowCustomerCount = '120px',
            userFollowCount = '100px',
            weworkChatCount = '100px',
            callCount = '100px',
            callRate = '100px',
            callDuration = '100px',
            smsCount = '100px',
            paidCount = '100px',
            orderCount = '100px',
            paidAmount = '130px'
        }
    }
}