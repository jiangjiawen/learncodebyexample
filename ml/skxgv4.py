import numpy as np
import xgboost as xgb
from sklearn.model_selection import StratifiedKFold
from sklearn.model_selection import learning_curve
from sklearn.model_selection import train_test_split,GridSearchCV
from sklearn.feature_selection import SelectFromModel
from joblib import Memory
from sklearn.datasets import load_svmlight_file
mem = Memory("./mycache")

@mem.cache
def get_data():
    data = load_svmlight_file("./data.txt")
    return data[0], data[1]

X, y = get_data()
X_train, X_test, y_train, y_test = train_test_split(
    X,
    y,
    train_size=0.8,
    test_size=0.2,
)

cls = xgb.XGBClassifier(
    **{
        "base_score": 0.5,
        "booster": "gbtree",
        "colsample_bylevel": 0.6,
        "colsample_bytree": 0.8,
        "gamma": 0.0,
        "learning_rate": 0.01,
        "max_delta_step": 0.0,
        "max_depth": 10,
        "min_child_weight": 5.0,
        "missing": None,
        "n_estimators": 500,
        "n_jobs": -1,
        "nthread": -1,
        "objective": "multi:softprob",
        "random_state": 7,
        "reg_alpha": 1.7,
        "reg_lambda": 1.2,
        "scale_pos_weight": 1,
        "seed": 10,
        "silent": True,
        "subsample": 0.6,
    }
)

model = SelectFromModel(cls)
model.fit(X_train, y_train)
print(X.shape)
X_new = model.transform(X_train)
X_test_new = model.transform(X_test)
feature_names = np.linspace(0,22,23,dtype=int)
print(feature_names)
print(feature_names[model.get_support()])
print(X_new.shape)

notconfirmeddata=load_svmlight_file("./testunknowndata.txt")
x_,y_=notconfirmeddata[0],notconfirmeddata[1]
x_selected = model.transform(x_)

model = xgb.XGBClassifier()
param_grid = {
    'booster':['gbtree'],
    'n_estimators': [500],
    'learning_rate': [0.01],
    'colsample_bytree': [0.7],
    'max_depth': [10],
    'reg_alpha': [1.7],
    'reg_lambda': [1.2],
    'subsample': [0.6]
}
kfold = StratifiedKFold(n_splits=3, shuffle=True, random_state=7)
grid_search = GridSearchCV(model, param_grid, scoring="neg_log_loss", n_jobs=-1, cv=kfold)
grid_result = grid_search.fit(X_new, y_train, verbose=1)

# summarize results
print(); print("Best: %f using %s" % (grid_result.best_score_, grid_result.best_params_))
# means = grid_result.cv_results_['mean_test_score']
# stds = grid_result.cv_results_['std_test_score']
# params = grid_result.cv_results_['params']
# for mean, stdev, param in zip(means, stds, params):
#     print("%f (%f) with: %r" % (mean, stdev, param))

from sklearn.metrics import classification_report,confusion_matrix
y_test_pred = grid_search.predict(X_test_new)
target_names=['0','1','2']
print(classification_report(y_test,y_test_pred,target_names=target_names))
print("confusion matrix:\n")
print(confusion_matrix(y_test,y_test_pred,labels=[0,1,2]))


y__pred = grid_search.predict(x_selected)
count_zero=0
count_one = 0 
count_tow=0
print("\n未确诊结果：")
for i in range(len(y__pred)):
    print(str(y__pred[i]),end=" ")
    if y__pred[i]==0:
        count_zero += 1
    if y__pred[i]==1:
        count_one += 1
    if y__pred[i]==2:
        count_tow += 1
print("\n\n数目0，1，2:")
print(count_zero,count_one,count_tow)