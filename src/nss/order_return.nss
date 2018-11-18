void OrderReturn(string sPSMessage)
{
  // -- this is returned from order to remove the busy variable from module
  if(sPSMessage == "uuid")
  {
    SetLocalInt(GetModule(),"uuidinprogress",0);
    WriteTimestampedLogEntry("Pubsub Heartbeat Event: "+sPSMessage);
  }
}