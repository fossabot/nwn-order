#include "nwnx_redis_ps"


void main()
{
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();

  if(data.channel == "heartbeat" && data.message == "1minute")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat" && data.message == "5minute")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat" && data.message == "30minute")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat"&& data.message == "1hour")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat" && data.message == "6hour")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat"&& data.message == "12hour")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat"&& data.message == "24hour")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }
}
