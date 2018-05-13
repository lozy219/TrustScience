import json
import numpy as np
from sklearn.model_selection import cross_val_score
from sklearn.linear_model import LogisticRegression
from sklearn import svm

f = open('nickname.json', 'r')
name2id = json.load(f)
f.close()

id2name = dict()
index = 0
for k in name2id:
    name2id[k] = index
    id2name[index] = k
    index += 1

data = list()
win = list()
f = open('match.csv', 'r')
for line in f:
    this_names = line.split()[0].split(',')
    new_data = np.zeros((1, index))
    for k in this_names[:5]:
        new_data[0, name2id[k]] += 1
    for k in this_names[5:]:
        new_data[0, name2id[k]] -= 1
    data.append(new_data)
    data.append(-new_data)
    win.append(np.ones((1, )))
    win.append(-np.ones((1, )))
f.close()

data = np.concatenate(data, axis=0)
win = np.concatenate(win, axis=0)

# cross validation code
# for C in [1e-10, 1e-8, 1e-6, 1e-4, 1e-2, 1e0, 1e2]:
#     model = LogisticRegression(C=C)
#     model = svm.SVC(C=C, kernel='linear')
#     scores = cross_val_score(model, data, win, cv=data.shape[0] // 2)
#     print(C, sum(scores) / len(scores))

model = LogisticRegression(C=1e-10)
model.fit(data, win)

scores = dict()
for i in range(len(id2name)):
    scores[id2name[i]] = int(model.coef_[0][i] * 1e11)

print(scores)
f = open('scores.json', 'w')
json.dump(scores, f, ensure_ascii=False)
f.close()
