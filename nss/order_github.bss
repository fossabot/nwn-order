
// -- this is called anytime order detects a webhook from github
void OrderGithub(string sWebhookType){
  if(sWebhookType == "commit")
    {
        WriteTimestampedLogEntry("Webhook recieved from github");
        WriteTimestampedLogEntry("Pubsub Heartbeat Event: "+sWebhookType);
    }
}