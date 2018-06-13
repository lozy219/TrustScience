# Author: 大队委

library('caret')
library('randomForest')
library('rpart')

match = read.csv('../data/match.csv',stringsAsFactors = F,encoding = 'UTF-8',header = F)
stat = read.csv('../data/stat.csv',stringsAsFactors = F,encoding = 'UTF-8',header = F)
colnames(stat) = c('shishen','win','lose','total','win_rate')
stat$match_rate = stat$total/378*100
stat$win_rate = as.integer(gsub("%","",stat$win_rate))
stat = stat[,c('shishen','win_rate','match_rate')]

result = data.frame()
for(i in 1:nrow(match)){
  
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

outcome = c(rep(1,189),rep(-1,189))
result[190:378,] = -result[190:378,]
dat = cbind(result,outcome)

summary(lm(outcome~.,data = dat))