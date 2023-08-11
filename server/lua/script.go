package lua

import "github.com/go-redis/redis"

var BuyItem = redis.NewScript(`
local amount = tonumber(ARGV[1])
local total = tonumber(redis.call("GET", "goodsTotal"))
local sold = tonumber(redis.call("GET", "goodsSold"))

if amount+sold>total then
    return 0
end

redis.call("INCRBY", "goodsSold", amount)

return amount
`)
