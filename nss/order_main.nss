#include "nwnx_redis"
#include "nwnx_redis_ps"

#include "order_return"
#include "order_github"
#include "order_heartbeat"

void main()
{
  // -- this is triggered via nwnxee redis pubusb. Do not trigger elsewhere.
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();
  
  // -- return function triggers
  if(data.channel == "return")
  {
      OrderReturn(data.message);
  }

  // -- return function for a github webhook being accepted.
  if(data.channel == "github")
  {
      OrderGithub(data.message);
  }

  // -- heartbeat functions
  if(data.channel == "heartbeat")
  {
      OrderHeartbeat(data.message);
  }
}