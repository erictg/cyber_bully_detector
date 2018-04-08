import pandas as pd
from keras.preprocessing import text
import cleaner
import matplotlib.pyplot as plt
import numpy as np
from keras import utils
from keras.layers import Dense, Embedding, Dropout, Activation, Conv1D, GlobalMaxPooling1D
from keras.models import Sequential
from sklearn.metrics import confusion_matrix
from sklearn.preprocessing import LabelEncoder
from keras.models import load_model


def tokenize_input(train, test, vocab):
    # tokenize input text
    tokenizer = text.Tokenizer(num_words=vocab)
    tokenizer.fit_on_texts(train)
    # create input matrix
    return tokenizer.texts_to_matrix(train), \
           tokenizer.texts_to_matrix(test), \
           tokenizer


def interactive_prompt(model, tokenizer, safe_word="Stop!"):
    """
    An interactive prompt for the cyber-bullying model
    :param model: Model object, should be compiled and trained
    :param tokenizer: Object used to tokenize training set
    :param safe_word: Word (or phrase) to exit prompt safely
    """
    print("Welcome to the cyber-bullying interactive prompt. Safe word is currently '" + safe_word + "'")
    while True:
        phrase = input("Enter a message to judge:")
        input_tokenized = tokenizer.texts_to_matrix([cleaner.clean_message(phrase)])
        prediction = model.predict(input_tokenized)
        print("Predicted label:", prediction[0])
        print("That's definitely an insult!" if prediction[0][0] < prediction[0][1]
              else "I don't think that's an insult!")


def init_data(filename, vocab_size, data_split):
    data = load_dataset(filename)
    train_input, test_input, train_label, test_label = split_dataset(data, data_split, "Comment", "Insult")
    x_train, x_test, tokenizer = tokenize_input(train_input, test_input, vocab_size)
    y_train, y_test, text_labels, num_classes = encode_labels(train_label, test_label)
    return train_input, test_input, train_label, test_label, x_train, x_test, y_train, y_test, tokenizer, text_labels,\
           num_classes


def load_dataset(filename):
    return pd.read_csv(filename, quotechar='|')


def load_modelfile(filename):
    return load_model(filename)


def split_dataset(dataset, split, input_field, label_field, ):
    train_size = int(len(dataset) * split)
    return dataset[input_field][:train_size], dataset[input_field][train_size:], \
           dataset[label_field][:train_size], dataset[label_field][train_size:]


def encode_labels(train, test):
    encoder = LabelEncoder()
    encoder.fit(train)
    y_train = encoder.transform(train)
    num_classes = np.max(y_train) + 1
    return utils.to_categorical(y_train, num_classes), \
           utils.to_categorical(encoder.transform(test), num_classes), \
           encoder.classes_, num_classes
