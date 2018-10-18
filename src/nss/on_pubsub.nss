#include "nwnx_redis_ps"
#include "nwnx_webhook"
#include "nwnx_admin"

void WebHookGo(string sReason,string sMessage);
void WebHookGo(string sReason,string sMessage)
  // ---- Webhook Data
  // -- Pull the webhook url data in module vars
  string sPublicWebhookUrl = GetLocalString(oMod,"PUBLIC_WEBHOOK");
  string sPrivateWebhookUrl = GetLocalString(oMod,"PRIVATE_WEBHOOK");

  // -- Timestamp
  string sCurrentTime = NWNX_Time_GetSystemTime();
  string sCurrentDate = NWNX_Time_GetSystemDate();

  // ---- Update Messages
  // -- Module update alert
  if(sReason == module && sMessage == update){
    string sWebhookMessage = "Module update found:" + sMessage + " || on: " + sCurrentDate;
  }  
  // -- Module Update proceed
  if(sReason == "module"){
    string sWebhookMessage = "Server rebooting for module update:" + sMessage + " || on: " + sCurrentDate;
  }
  // -- Nwnxee Update proceed
  if(sReason == "nwnxee"){
    string sWebhookMessage = "Server rebooting for nwnxee update:" + sMessage + " || on: " + sCurrentDate;
  }

  // -- actually send the message
  NWNX_WebHook_SendWebHookHTTPS("discordapp.com",sPublicWebhookUrl, sWebhookMessage , "Module Update");

// -- The fun stuff
void ContinuousIntegration(string sReason, string sMessage);
void ContinuousIntegration(string sReason, string sMessage)
{    
  // -- Module update alert
  if(sReason == module && sMessage == update){
    string sBuiltstring = "Module update detected: " + sMEssage;
    ActionSpeakString(sBuiltstring,TALKVOLUME_SHOUT);
    WebHookGo(sReason,sMessage)
    break;
  }  
  
  // ---- Shutdown process
  object oMod = GetModule();

  // -- Nwnxee Character save process
  if(sReason == "nwnxee"){
    FadeToBlack(oPC,FADE_SPEED_FAST);
    ExportSingleCharacter(oPC);
    BootPC(oPC,"Server Rebooting -- Character has been saved");
    oPC = GetNextPC();
  }

  // -- Module update save character just cause
  if(sReason == "module"){
    ExportSingleCharacter(oPC);
    oPC = GetNextPC();

  }
  
  // -- Password lock the server
  NWNX_Administration_SetPlayerPassword("!@#$%^&*()");
  
  // -- Save the database
  NWNX_Redis_SAVE();

  // -- Webhooks
  WebHookGo(sReason,sMessage)

  // -- If base nwnxee docker image updates we have to restart
  if(sReason == "nwnxee"){
    NWNX_Administration_ShutdownServer();
  }

  // -- If module is updated no need to restart
  if(sReason == "module"){
    StartNewModule(GetName(oMod));
  }
}

void main()
{
  struct NWNX_Redis_PubSubMessageData data = NWNX_Redis_GetPubSubMessageData();

  // -- Triggers when pubsub message comes in on the "updating" channel
  if(data.channel == "module" && data.message == "alert")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
    ContinuousIntegration(data.channel,data.message);
  }

  // -- Triggers when pubsub message comes in on the "module" channel
  if(data.channel == "module"&& data.message == "proceed")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
    ContinuousIntegration(data.channel,data.message);
  }

  // -- Triggers when pubsub message comes in on the "nwnxee" channel
  if(data.channel == "nwnxee" && data.message == "proceed")
  {
    WriteTimestampedLogEntry("Pubsub Event: channel=" + data.channel + " message=" + data.message);
    ContinuousIntegration(data.channel,data.message);
  }
}
