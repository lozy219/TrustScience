# Author: 大队委

library('caret')
library('randomForest')
library('rpart')

#setwd("~/Documents/TrustScience")
match = read.csv('../data/match.csv',stringsAsFactors = F,encoding = 'UTF-8',header = F)
stat = read.csv('../data/stat.csv',stringsAsFactors = F,encoding = 'UTF-8',header = F)
colnames(stat) = c('shishen','win','lose','total','win_rate')
stat$match_rate = stat$total/nrow(match)*100
stat$win_rate = as.integer(gsub("%","",stat$win_rate))
stat = stat[,c('shishen','win_rate','match_rate')]

train_index = sample(nrow(match),0.7*nrow(match))
match_train = match[train_index,]
match_test = match[-train_index,]

result = data.frame()
for(i in 1:nrow(match_train)){
  
  win_rate1 = stat$win_rate[stat$shishen == match$V1[i]]
  win_rate2 = stat$win_rate[stat$shishen == match$V2[i]]
  win_rate3 = stat$win_rate[stat$shishen == match$V3[i]]
  win_rate4 = stat$win_rate[stat$shishen == match$V4[i]]
  win_rate5 = stat$win_rate[stat$shishen == match$V5[i]]
  
  match_rate1 = stat$match_rate[stat$shishen == match$V1[i]]
  match_rate2 = stat$match_rate[stat$shishen == match$V2[i]]
  match_rate3 = stat$match_rate[stat$shishen == match$V3[i]]
  match_rate4 = stat$match_rate[stat$shishen == match$V4[i]]
  match_rate5 = stat$match_rate[stat$shishen == match$V5[i]]
  
  win_rate6 = stat$win_rate[stat$shishen == match$V6[i]]
  win_rate7 = stat$win_rate[stat$shishen == match$V7[i]]
  win_rate8 = stat$win_rate[stat$shishen == match$V8[i]]
  win_rate9 = stat$win_rate[stat$shishen == match$V9[i]]
  win_rate10 = stat$win_rate[stat$shishen == match$V10[i]]
  
  match_rate6 = stat$match_rate[stat$shishen == match$V6[i]]
  match_rate7 = stat$match_rate[stat$shishen == match$V7[i]]
  match_rate8 = stat$match_rate[stat$shishen == match$V8[i]]
  match_rate9 = stat$match_rate[stat$shishen == match$V9[i]]
  match_rate10 = stat$match_rate[stat$shishen == match$V10[i]]
  
  win_win_rate = mean(c(win_rate1,win_rate2,win_rate3,win_rate4,win_rate5))
  win_match_rate = mean(c(match_rate1,match_rate2,match_rate3,match_rate4,match_rate5))
  
  lose_win_rate = mean(c(win_rate6,win_rate7,win_rate8,win_rate9,win_rate10))
  lose_match_rate = mean(c(match_rate6,match_rate7,match_rate8,match_rate9,match_rate10))
  
  result = rbind(result,c(win_win_rate-lose_win_rate,win_match_rate-lose_match_rate))
}
colnames(result) = c('win_rate.diff','match_rate.diff')

half = floor(nrow(match_train)/2)
outcome = c(rep(1,half),rep(-1,nrow(match_train)-half))
result[(half+1):nrow(match_train),] = -result[(half+1):nrow(match_train),]
dat = cbind(result,outcome)

## linear regression
model.lm = lm(outcome~win_rate.diff+match_rate.diff, data = dat)
summary(model.lm)
stat$score_lm = coefficients(model.lm)[2] * stat$win_rate + coefficients(model.lm)[3] * stat$match_rate


## logistic regression
dat$outcome_glm = as.factor(dat$outcome)
model.glm = glm(outcome_glm~win_rate.diff+match_rate.diff, data = dat, family = 'binomial')
summary(model.glm)
stat$score_glm = coefficients(model.glm)[2] * stat$win_rate + coefficients(model.glm)[3] * stat$match_rate




for(i in 1:nrow(match)){
  team1_score_lm = mean(sapply(match[i,1:5],function(x){stat$score_lm[stat$shishen == x]}))
  team2_score_lm = mean(sapply(match[i,6:10],function(x){stat$score_lm[stat$shishen == x]}))
  team1_score_glm = mean(sapply(match[i,1:5],function(x){stat$score_glm[stat$shishen == x]}))
  team2_score_glm = mean(sapply(match[i,6:10],function(x){stat$score_glm[stat$shishen == x]}))
  
  match$score_diff_lm[i] = team1_score_lm - team2_score_lm + coefficients(model.lm)[1]
  match$score_diff_glm[i] = exp(team1_score_glm - team2_score_glm+coefficients(model.glm)[1])/(1+exp(team1_score_glm - team2_score_glm+coefficients(model.glm)[1]))
}


for(threshold in seq(0,0.6,0.1)){
  train_match_num = sum(match$score_diff_lm[train_index] >= threshold | match$score_diff_lm[train_index] <= -threshold)
  train_accuracy = sum(match$score_diff_lm[train_index] >= threshold)/train_match_num*100
  test_match_num = sum(match$score_diff_lm[-train_index] >= threshold | match$score_diff_lm[-train_index] <= -threshold)
  test_accuracy = sum(match$score_diff_lm[-train_index] >= threshold)/test_match_num*100
  cat("threshold:",threshold,"  train_accuracy:",train_accuracy,'  train_match_num:',train_match_num,'\n')
  cat("threshold:",threshold,"  test_accuracy:",test_accuracy,'  test_match_num:',test_match_num,'\n')
}

for(threshold in seq(0.4,0.6,0.05)){
  train_accuracy = sum(match$score_diff_glm[train_index] >= threshold)/nrow(match_train)*100
  test_accuracy = sum(match$score_diff_glm[-train_index] >= threshold)/nrow(match_test)*100
  cat("threshold:",threshold,"  train_accuracy:",train_accuracy,'  test_accuracy:',test_accuracy,'\n')
}



## Visualization - shishen match_rate & win_rate
plot(x = stat$win_rate, y = stat$match_rate,xlim = c(20,100),ylim = c(0,30),xlab = 'win_rate',ylab = 'match_rate',cex = 0.1)
text(stat$win_rate, stat$match_rate, labels = stat$shishen,family = "Heiti SC Light", cex = 0.6, pos = 3, col = 'blue')
grid(NULL,NULL)


## Visualization - logistic regression
