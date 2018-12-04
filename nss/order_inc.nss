#include "nwnx_redis"

#include "order_func"
#include "order_github"
#include "order_heartbeat"
#include "order_return"

#include "_save"
#include "_log"

string RdsEdgePlayer(string sEdgeType,object oPC);
string RdsEdgePlayer(string sEdgeType,object oPC) {
  string Nwserver = GetModuleName();
  string CDKey    = GetPCPublicCDKey(oPC, FALSE);
  string UUID     = OrderGetUUIDPlayer(oPC);
  if (sEdgeType == "player") {
    string sEdge = Nwserver+":player:"+UUID;
    return sEdge;
  } else {
    return "error";
  }
}


string RdsEdgeServer(string sEdgeType);
string RdsEdgeServer(string sEdgeType) {
  string Nwserver = GetModuleName();
  if (sEdgeType == "server") {
    string sServerEdge = Nwserver+":server:";
    return sServerEdge;
  }
  else if (sEdgeType == "item") {
    string sItemsEdge = Nwserver+":item:";
    return sItemsEdge;
  } else {
    return "error";
  }
}