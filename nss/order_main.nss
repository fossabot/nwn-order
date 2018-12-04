#include "nwnx_redis_ps"

#include "order_inc"

void main() {
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();
  Log(data.channel, "1");
  int pubusbType = StringToInt(data.channel);

  if (data.channel == "input") {
    OrderReturn(data.message);
  } else if (data.channel == "github") {
    OrderGithub(data.message);
  } else if (data.channel == "heartbeat") {
    OrderHeartbeat(data.message);
  } else {
    Log("Error, channel type not recognized.","1");
  }
}