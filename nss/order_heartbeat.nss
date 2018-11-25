// -- this is what is triggered via the order heartbeat tickers.
void OrderHeartbeat(string sTicker){
    int nTicker = StringToInt(sTicker);

    switch (nTicker) 
    {
        // one minute heartbeat
        case 1 :{
            // stuff
            break;
        }
        // five minute heartbeat
        case 5 :{
            // stuff
            break;
        }
        // thirty minute heartbeat
        case 30 :{
            // stuff
            break;
        }
        // one hour heartbeat
        case 60 :{
            // stuff
            break;
        }
        // six hour heartbeat
        case 360:{
            // stuff
            break;
        }
        // twelve hour heartbeat
        case 720 :{
            // stuff
            break;
        }
        // twentyfour hour heartbeat
        case 1440:{
            // -- examples
            // SaveModule();
            // RestartModule();
            break;
        }
    }
}