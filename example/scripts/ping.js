#!js name=function api_version=1.0

function ping(client, data) {
    return client.call("ping");
}

redis.registerFunction("ping", ping);