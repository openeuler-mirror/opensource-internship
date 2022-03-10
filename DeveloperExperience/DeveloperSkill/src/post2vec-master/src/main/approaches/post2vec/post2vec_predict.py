import sys
sys.path.append('./../../../')
from torch.autograd import Variable
import torch.nn.functional as F
import os
import sys
import torch
import torch.nn as nn
from nltk import word_tokenize
from utils.eval_util import evaluate_batch, evaluate_batch_f1_5
import numpy as np
from utils.data_util import get_specific_comp_list
from pathConfig import data_dir
from utils.data_util import random_mini_batch, load_pickle
import os
from utils.padding_and_indexing_util import padding_and_indexing_qlist
from utils.vocab_util import vocab_to_index_dict
from main.approaches.post2vec.post2vec_util import load_args, load_model
from utils.file_util import read_file_str_list, write_str_to_file
from utils.time_util import get_current_time
from data_structure.question import Question

def predict(title, desc_text, desc_code):
    task = 'tagRec'
    dataset = "SO-05-Sep-2018"
    dataset_dir = data_dir + os.sep + task + os.sep + dataset
    # ts dir
    ts = 50
    ts_dir = dataset_dir + os.sep + "ts%s" % ts
    # sample_K dir
    sample_K = "test100000"
    sample_K_dir = ts_dir + os.sep + "data-%s" % sample_K
    vocab_dir = os.path.join(sample_K_dir, "vocab")
    app_dir = os.path.join(sample_K_dir, "approach", "post2vec")
    snapshot_dir = os.path.join(app_dir, "snapshot")

    # basic path
    print("Setting:\ntask : %s\ndataset : %s\nts : %s\n" % (task, dataset, ts))
    #################################################################################

    # load vocab
    # initial
    len_dict_fpath = os.path.join(vocab_dir, "len.pkl")
    title_vocab_fpath = os.path.join(vocab_dir, "title_vocab.pkl")
    desc_text_vocab_fpath = os.path.join(vocab_dir, "desc_text_vocab.pkl")
    desc_code_vocab_fpath = os.path.join(vocab_dir, "desc_code_vocab.pkl")
    tag_vocab_fpath = os.path.join(vocab_dir, "tag_vocab.pkl")

    # len
    # len_dict = load_pickle(len_dict_fpath)
    len_dict = dict()
    len_dict["max_title_len"] = 100
    len_dict["max_desc_text_len"] = 1000
    len_dict["max_desc_code_len"] = 1000

    # title vocab
    title_vocab = load_pickle(title_vocab_fpath)
    title_vocab = vocab_to_index_dict(vocab=title_vocab, ifpad=True)

    # desc_text vocab
    desc_text_vocab = load_pickle(desc_text_vocab_fpath)
    desc_text_vocab = vocab_to_index_dict(vocab=desc_text_vocab, ifpad=True)

    # desc_code_vocab
    desc_code_vocab = load_pickle(desc_code_vocab_fpath)
    desc_code_vocab = vocab_to_index_dict(vocab=desc_code_vocab, ifpad=True)

    # tag vocab
    tag_vocab = load_pickle(tag_vocab_fpath)
    tag_vocab = vocab_to_index_dict(vocab=tag_vocab, ifpad=False)
    vocab_list = load_pickle(tag_vocab_fpath)
    vocab_list = list(vocab_list)
    vocab_list.sort()
    # predict
    test_dir = os.path.join(sample_K_dir, "test")

    # load args from json file
    snapshot_dir_name = "2022-03-09_08-06-22"
    param_dir = os.path.join(snapshot_dir, snapshot_dir_name)
    args = load_args(param_dir)

    topk_list = [1, 2, 3, 4, 5]

    res_str = ''
    param_name = 'snapshot_steps_160500.pt'
    param_fpath = os.path.join(param_dir, param_name)

    model = load_model(args, param_fpath)

    if args.model_selection == "all":
        from main.approaches.post2vec.models.model_all import eval
    elif args.model_selection == "title":
        from main.approaches.post2vec.models.model_title import eval
    elif args.model_selection == "title_desc_text":
        from main.approaches.post2vec.models.model_title_desc_text import eval

    tokenized_title = word_tokenize(title.lower())
    tokenized_desc_text = word_tokenize(desc_text.lower())
    tokenized_desc_code = word_tokenize(desc_code.lower())
    predict_data = Question(0, tokenized_title, tokenized_desc_text, tokenized_desc_code, 0, "tags")
    print(predict_data.title, predict_data.desc_text, predict_data.desc_code)
    predict_datas = [predict_data]

    processed_predict_datas = padding_and_indexing_qlist(predict_datas, len_dict, title_vocab, desc_text_vocab,
                                                    desc_code_vocab, tag_vocab)

    with torch.no_grad():
        model.eval()
        # features
        t = get_specific_comp_list("title", processed_predict_datas)
        dt = get_specific_comp_list("desc_text", processed_predict_datas)
        dc = get_specific_comp_list("desc_code", processed_predict_datas)
        # label
        target = np.array(get_specific_comp_list("tags", processed_predict_datas))

        t = torch.tensor(t).long()
        dt = torch.tensor(dt).long()
        dc = torch.tensor(dc).long()
        target = torch.tensor(target).float()

        if args.cuda:
            t, dt, dc, target = t.cuda(), dt.cuda(), dc.cuda(), target.cuda()

        logit = model(t, dt, dc)
        if torch.cuda.is_available():
            preds = logit.cpu().detach().numpy()
        else:
            preds = logit.detach().numpy()
        topk = 20
        for pred in preds:
            top_idx_list = sorted(range(len(pred)), key=lambda i: pred[i])[-topk:]
            for index in range(len(top_idx_list) - 1, -1, -1):
                print(vocab_list[top_idx_list[index]])


if __name__ == '__main__':
    title = "openEuler2203】ERROR: modpost: \"show\" [/home/cc/cc/mod2.ko] undefined!"
    desc_text = "Problem Description: Is it not supported that a kernel module depends on another kernel module to be compiled? Edit two kernel modules mod1 and mod2. The kernel module mod2 calls a function of the module mod1. When compiling, I first compile the module mod1. The module mod1 will export the corresponding function to mod2 for use. Use the Module.symvers of the module mod1 to compile the module. Report undefined when mod2"
    desc_code = "#include <linux/init.h> #include <linux/module.h> int global_var = 100; void show(void) { printk(\"show(): global_var =%d \n\",global_var); } static int hello_init(void) { printk(\"module b :global_var=%d\n\",global_var); return 0; } static void hello_exit(void) { printk(\"hello_exit \n\"); return; } EXPORT_SYMBOL(global_var); EXPORT_SYMBOL(show); MODULE_AUTHOR(\"yikoulinux\"); MODULE_LICENSE(\"GPL\"); module_init(hello_init); module_exit(hello_exit);"
    predict(title, desc_text, desc_code)