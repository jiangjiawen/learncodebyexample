import numpy as np
from sklearn.svm import SVC
from sklearn.preprocessing import StandardScaler
from sklearn.model_selection import GridSearchCV, train_test_split
from sklearn.feature_selection import SelectFromModel
from joblib import Memory
from sklearn.datasets import load_svmlight_file
mem = Memory("./mycache")

@mem.cache
def get_data():
    data = load_svmlight_file("./data.txt")
    return data[0], data[1]

X, y = get_data()
# svc feature_selection 有点小问题 于是使用lsvc选择feature
from sklearn.svm import LinearSVC
lsvc = LinearSVC(C=0.01, penalty="l1", dual=False).fit(X, y)
# svc = SVC(kernel='rbf', class_weight='balanced',).fit(X,y)
model = SelectFromModel(lsvc,prefit=True)
# model.fit(X, y)
print(X.shape)
X_new = model.transform(X)
print(X_new.shape)
x_train, x_test, y_train, y_test = train_test_split(X_new, y, test_size=.2)

svc = SVC(kernel='rbf', class_weight='balanced',)
c_range = np.logspace(-2, 3, 20, base=2)
gamma_range = np.logspace(-1, 4, 20, base=2)
param_grid = [{'kernel': ['rbf'], 'C': c_range, 'gamma': gamma_range}]
grid = GridSearchCV(svc, param_grid, cv=3, n_jobs=-1)
clf = grid.fit(x_train, y_train)
print(grid.best_params_)
score = grid.score(x_test, y_test)
print('准确度为%s' % score)

from sklearn.metrics import classification_report
y_test_pred = grid.predict(x_test)
target_names=['0','1','2']
print(classification_report(y_test,y_test_pred,target_names=target_names))