#include "nwnx_redis"

// -- return or assign and return the oPC uuid.
string OrderGetUUIDPlayer(object oPC);
string OrderGetUUIDPlayer(object oPC) {
  // if the user has no uuid set
  if (GetTag(oPC) == "") {  
    object oMod = GetModule();
        
    // -- Confirm we aren't conflicting keys.
    int nUuidInProgress = GetLocalInt(oMod,"uuidinprogress"); 

    // -- If in progress, else return ""                                                                                       
    if (nUuidInProgress != 1) {
      // get prepared uuid
      int nUUID = NWNX_Redis_GET("system:uuid");
      string sUUID = NWNX_Redis_GetResultAsString(nUUID);
      WriteTimestampedLogEntry(sUUID);
      SetTag(oPC, sUUID);
      // delete the key after we get the value set to sUUID
      NWNX_Redis_DEL("system:uuid"); 
      // send out task to generate new uuid via order.
      NWNX_Redis_PUBLISH("input","newuuid");
      // sets the module variable so we don't end up duplicating anything.
      SetLocalInt(oMod,"uuidinprogress",1);

      return sUUID;
    } else {
      // if no uuid can be grabbed, return "", which should be filtered from being saved. 
      return "";
    }
  } else {
      string sUUID = GetTag(oPC);
      return sUUID;
  }
}

// -- return a uuid unrelated to player
string OrderGetUUID();
string OrderGetUUID() {
  object oMod = GetModule();
  // get cached uuid
  int nUUID = NWNX_Redis_GET("system:uuid");
  string sUUID = NWNX_Redis_GetResultAsString(nUUID);
  // delete the key after we get the value set to sUUID
  NWNX_Redis_DEL("system:uuid"); 
  // send out task to generate new uuid via order.
  NWNX_Redis_PUBLISH("input","newuuid");
  // sets the module variable so we don't end up duplicating anything.
  SetLocalInt(oMod,"uuidinprogress",1);
  // return a uuid
  return sUUID;
}