A. SMS Spam Collection v.1
-------------------------

### 1. Description

The SMS Spam Collection v.1 (hereafter the corpus) is a set of SMS tagged messages that have been collected for SMS Spam research. It contains one set of SMS messages in English of 4592 messages, tagged acording being ham (legitimate) or spam. 

### 2. Format

The files contain one message per line. Each line is composed by two columns: one with label (ham or spam) and other with the raw text. Here are some examples:

ham   What you doing?how are you?  
ham   Ok lar... Joking wif u oni...  
spam   FreeMsg: Txt: CALL to No: 86888 & claim your reward of 3 hours talk time to use from your phone now! ubscribe6GBP/ mnth inc 3hrs 16 stop?txtStop  
spam   Sunshine Quiz! Win a super Sony DVD recorder if you canname the capital of Australia? Text MQUIZ to 82277. B

Note: messages are not chronologically sorted.

### 3. License/Disclaimer

Read more at http://www.dt.fee.unicamp.br/~tiago/smsspamcollection/ 


B. SMS Test
-----------

### 1. Description

The SMS Test is a set of SMS that have been separated from **SMS Spam Collection**. It contains the last 980 messages.

### 2. Format

The files contain one message per line. All messages used for testing Bayesian classifier which written in **bayes.go**
