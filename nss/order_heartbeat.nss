#include "_log"

// -- this is what is triggered via the order heartbeat tickers.
void OrderHeartbeat(string sTicker) {
    Log("heartbeat: "+sTicker,"1");
    if (sTicker == "1") {

    } else if (sTicker == "5") {

    } else if (sTicker == "30") {

    } else if (sTicker == "60") {

    } else if (sTicker == "360") {

    } else if (sTicker == "720") {

    } else if (sTicker == "1440") {

    } else {
      Log("Error, ticker not recognized.","1");
    }
}