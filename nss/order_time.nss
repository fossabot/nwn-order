#include "order_main"

int OrderGetTimeHour(); 
int OrderGetTimeHour(){
    string sTime = NWNX_Redis_HMGET(RdsPlayerEdge(RdsServerEdge()+":serverstats", "ingametime:hour");
    int nTime = StringToInt(sTime);
    return nTime;
}

int OrderGetTimeMinute(); 
int OrderGetTimeMinute(){
    string sTime = NWNX_Redis_HMGET(RdsPlayerEdge(RdsServerEdge()+":serverstats", "ingametime:minute");
    int nTime = StringToInt(sTime);
    return nTime;
}

int OrderGetTimeSecond(); 
int OrderGetTimeSecond(){
    string sTime = NWNX_Redis_HMGET(RdsPlayerEdge(RdsServerEdge()+":serverstats", "ingametime:second");
    int nTime = StringToInt(sTime);
    return nTime;
}

int OrderGetTimeMillisecond();
int OrderGetTimeMillisecond(){
    string sTime = NWNX_Redis_HMGET(RdsPlayerEdge(RdsServerEdge()+":serverstats", "ingametime:millisecond");
    int nTime = StringToInt(sTime);
    return nTime;
} 