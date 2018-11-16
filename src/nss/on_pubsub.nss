#include "nwnx_redis_ps"
#include "nwnx_redis"

string uuid();
string uuid(){
    object oMod = GetModule();

    // get cached uuid
    string sUUID = NWNX_Redis_GET("system:uuid");

    // delete the key after we get the value set to sUUID
    NWNX_Redis_DEL("system:uuid"); 
    // send out task to generate new uuid via order.
    NWNX_Redis_PUBLISH("input","newuuid");
    SetLocalInt(oMod,"uuidinprogress",1);

    return sUUID;
}

void main()
{
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();

  if(data.channel == "return" && data.message == "uuid")
  {
    object oMod = GetModule();
    SetLocalInt(oMod,"uuidinprogress",0);
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=uuidinprogress variable set to 0");
  }

  if(data.channel == "heartbeat" && data.message == "1m")
  {
    // WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
    WriteTimestampedLogEntry(uuid());
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
  
  if(data.channel == "github"&& data.message == "repoupdate")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
  }

  if(data.channel == "uuid"&& data.message == "uuid")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
    
  }
}
