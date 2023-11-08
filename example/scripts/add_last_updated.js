#!js name=trigger api_version=1.0

function setLastUpdate(client, data) {
    if(data.event == "hset") {
        const time = new Date();
        client.call("hset", data.key, "last_update", time.toISOString());
    }
}

redis.registerKeySpaceTrigger("addLastUpdated", "shop:", setLastUpdate);