#include "nwnx_redis"


//string OrderPlayerUUID(object oPC);
//string OrderPlayerUUID(object oPC)
//{
//    if (GetTag(oPC) == ""){  
//        object oMod = GetModule();
//        // confirm we aren't stealing someone elses key.
//        int nUuidInProgress = GetLocalInt(oMod,"uuidinprogress","1");                                                                                         
//        while (nUuidInProgress != "1")
//        {
//            // get prepared uuid
//            string sUUID = NWNX_Redis_GET("system:uuid");
//            WriteTimestampedLogEntry(sUUID);
//            // delete the key after we get the value set to sUUID
//            NWNX_Redis_DEL("system:uuid"); 
//            // send out task to generate new uuid via order.
//            NWNX_Redis_PUBLISH("input","newuuid");
//            SetLocalInt(oMod,"uuidinprogress","1");
//            return sUUID;
//        }
//        else
//        {
//            DelayCommand(5.0f,OrderPlayerUUID)
//            break;
//        }
//    else{
//        OrderPlayerUUID(oPC);
//        return "noID";
//    }
//    }
//    return "";
//}   