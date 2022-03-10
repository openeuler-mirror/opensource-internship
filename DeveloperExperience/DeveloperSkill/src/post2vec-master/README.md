# Developer Skill Tag Recommendation Model

## Data Download

- Data download [here [12.8GB]](https://drive.google.com/file/d/1g1tAebVnT76pYcY74IxyqoMU7KBEXPmb/view?usp=sharing).
- Or you can download raw data from [Stack Overflow data dump](https://archive.org/download/stackexchange), the data preprocessing scripts are provided [here](https://github.com/post2vec/post2vec/tree/master/src/data_preparation).
- The trained model parameter file can be found [here](https://bhpan.buaa.edu.cn:443/link/FF43BAFB10DA709E6D8B80DCFDAEC07D)

## Experiment Instructions

| -       | Command                                                  |
| ------- | -------------------------------------------------------- |
| -       | Post2Vec                                                 |
| Train   | ```python src/approaches/post2vec/post2vec_train.py```   |
| Test    | ```python src/approaches/post2vec/post2vec_test.py```    |
| Predict | ```python src/approaches/post2vec/post2vec_predict.py``` |


## Experiemnt result

Due to the large amount of data, we only use 10% of the data for training, and the results of the reproduction experiment are as follows.

| Precision@1 | Precision@2 | Precision@3 | Precision@4 | Precision@5 |
| ----------- | ----------- | ----------- | ----------- | ----------- |
| 0.7079      | 0.5718      | 0.4651      | 0.3873      | 0.3312      |
| Recall@1    | Recall@2    | Recall@3    | Recall@4    | Recall@5    |
| 0.7079      | 0.6140      | 0.5812      | 0.5865      | 0.6073      |
| F1-score@1  | F1-score@2  | F1-score@3  | F1-score@4  | F1-score@5  |
| 0.7079      | 0.5859      | 0.5023      | 0.4480      | 0.4101      |