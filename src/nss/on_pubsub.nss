#include "nwnx_redis_ps"


void main()
{
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();

  if(data.channel == "heartbeat" && data.message == "1m")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat" && data.message == "5m")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat" && data.message == "30m")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat"&& data.message == "1h")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat" && data.message == "6h")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat"&& data.message == "12h")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "heartbeat"&& data.message == "24h")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }
}
