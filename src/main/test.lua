function selectMobiles(meetingIds, tagIds, mode)
    local result = {}
    if next(meetingIds) == nil
    then
        return result;
    end
    local sql = [[
        select user_id                as userId,
       user_name              as   userName,
       count(mcpr.id)         as ordercount,
       count(distinct msd.id) as invit,
       sum(mcpr.paid_amount)
from meeting_commerce_performance_record mcpr
         left join meeting_send_detail msd on mcpr.biz_id = msd.biz_id and mcpr.meeting_id = msd.meeting_id
where mcpr.meeting_id = 894
  and msd.send_status = 1;
    ]]

    local params = {
        meetingIds = string.join(",", meetingIds)
    }
    -- 标签ID筛选
    if tagIds ~= "" and tagIds ~= nil then
        -- 把 "a,b,c,d" 替换成 "'a','b','c','d'"，用于 in 查询条件
        local tagIds = "'" .. request.tagIds:gsub(",", "','") .. "'"

        local existTag = [[AND exists (
        select 1 from wework_contact_tag t
        where t.biz_id = mci.biz_id 
        and mci.contact_id != ''
        and t.contact_id = mci.contact_id
        and t.`type` = 1
        and t.tag_id in ({tagIds})
        and t.is_deleted = 0
        )]]
        sql = sql:gsub("{existTag}", existTag)
        params["tagIds"] = tagIds
    else
        sql = sql:gsub("{existTag}", "")
    end
     -- 群ID筛选
    if request.chatRoomIds ~= "" then
        -- 把 "a,b,c,d" 替换成 "'a','b','c','d'"，用于 in 查询条件
        local chatRoomIds = "'" .. request.chatRoomIds:gsub(",", "','") .. "'"

        local existChatRoom = [[AND exists (
        select 1 from wework_chat_room_relation r
        where r.biz_id = mci.biz_id 
        and mci.contact_id != ''
        and r.member_id = mci.contact_id
        and r.wework_room_id in ({chatRoomIds})
        and r.is_deleted = 0
        )]]
    
        sql = sql:gsub("{existChatRoom}", existChatRoom)
        params["chatRoomIds"] = chatRoomIds
    else
        sql = sql:gsub("{existChatRoom}", "")
    end

    local fetch_res = mysql_fetch_all('scrm', sql % params)
    local i = 1;
    if fetch_res then
        for k,v in pairs(fetch_res) do
            result[i] = v.mobile
            if mode == 1 then
                result[i] = v.num
            end
            i = i + 1
        end
    end
    return result;
end

function selectCount(mobiles)
    if next(mobiles) == nil
    then
        return 0;
    end
    local sql = [[
	SELECT count(*) count
    FROM
	    {schema_name}.leads_lib ll
    WHERE
	    ll.mobile IN ({mobiles})
	    or ll.mobile1 IN ({mobiles})
    ]]

    local params = {
        mobiles = string.join(",", mobiles)
    }
    sql = sql:gsub('{schema_name}', schema_name)
    local fetch_res = mysql_fetch_one('crm', sql % params)
    if fetch_res then 
        return fetch_res['count'];
    end
    return 0;
end

function selectUserCount(meetingIds, nums)
    if next(meetingIds) == nil or nums == nil
    then
        return 0;
    end
    local sql = [[
	SELECT count(DISTINCT
		meeting_customer_num ) count
	FROM
		meeting_customer_order_relation mcor
		LEFT JOIN shop_order so ON mcor.order_id = so.order_id
		LEFT JOIN meeting_customer_info mci ON mcor.meeting_customer_num = mci.num
	WHERE
		mcor.meeting_id IN ({meetingIds}) 
		and mci.num IN ({nums})
		AND so.`status` = 10000;
    ]]

    local params = {
        meetingIds = string.join(",", meetingIds),
        nums = "'" .. string.join("','", nums) .. "'"
    }
 
    sql = sql % params
    local fetch_res = mysql_fetch_one('scrm', sql)
    local result = fetch_res["count"]
 
    return result;
end

-- 获取meetingIds
request.startTime = ''
request.endTime = ''
local meetingList = getMeetingList()
local meetingIds = {}
local i = 1
for k, v in pairs(meetingList) do
    meetingIds[i] = v.id
    i = i + 1
end

local roomIds = request.chatRoomIds
local tagIds = request.tagIds

template = {
    data1 = '--',
    data2 = '--',
    data3 = '--'
}
if isOpenCrm() and i == 2 then
    local progFormat = "%.2f"
    local mobiles = selectMobiles(meetingIds, tagIds, 0)
    local nums = selectMobiles(meetingIds, tagIds, 1)

    local data1 = selectCount(mobiles)
    local data2 = selectUserCount(meetingIds, nums)
    local data3 = "--"
    if data1 ~= 0 then
        data3 = progFormat:format(data2 / data1 * 100.0) .. "%"
    end
    
    template = {
        data1 = data1,
        data2 = data2,
        data3 = data3
    }
end

return {
    template = template
}