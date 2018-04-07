import itertools

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from keras import utils
from keras.layers import Dense, Activation, Dropout
from keras.models import Sequential
from keras.preprocessing import text
from sklearn.metrics import confusion_matrix
from sklearn.preprocessing import LabelEncoder


# This utility function is from the sklearn docs:
# http://scikit-learn.org/stable/auto_examples/model_selection/plot_confusion_matrix.html
def plot_confusion_matrix(cm, classes, title='Confusion matrix', cmap=plt.cm.Blues):
    """
    This function prints and plots the confusion matrix.
    Normalization can be applied by setting `normalize=True`.
    """
    cm = cm.astype('float') / cm.sum(axis=1)[:, np.newaxis]

    plt.imshow(cm, interpolation='nearest', cmap=cmap)
    plt.title(title, fontsize=30)
    plt.colorbar()
    tick_marks = np.arange(len(classes))
    plt.xticks(tick_marks, classes, rotation=45, fontsize=22)
    plt.yticks(tick_marks, classes, fontsize=22)

    fmt = '.2f'
    thresh = cm.max() / 2.
    for i, j in itertools.product(range(cm.shape[0]), range(cm.shape[1])):
        plt.text(j, i, format(cm[i, j], fmt),
                 horizontalalignment="center",
                 color="white" if cm[i, j] > thresh else "black")

    plt.ylabel('True label', fontsize=25)
    plt.xlabel('Predicted label', fontsize=25)


def load_dataset(filename):
    return pd.read_csv(filename)


def split_dataset(dataset, split, input_field, label_field, ):
    train_size = int(len(dataset) * split)
    return dataset[input_field][:train_size], dataset[input_field][train_size:], \
           dataset[label_field][:train_size], dataset[label_field][train_size:]


def tokenize_input(train, test, vocab):
    # tokenize input text
    tokenize = text.Tokenizer(num_words=vocab)
    tokenize.fit_on_texts(train)
    # create input matrix
    return tokenize.texts_to_matrix(train),\
           tokenize.texts_to_matrix(test),\
           tokenize


def encode_labels(train, test):
    encoder = LabelEncoder()
    encoder.fit(train)
    y_train = encoder.transform(train)
    num_classes = np.max(y_train) + 1
    return utils.to_categorical(y_train, num_classes),\
           utils.to_categorical(encoder.transform(test), num_classes),\
           encoder.classes_, num_classes


def create_model(vocab_size, num_classes):
    # Build the model
    model = Sequential()
    model.add(Dense(512, input_shape=(vocab_size,)))
    model.add(Activation('relu'))
    model.add(Dropout(0.5))
    model.add(Dense(num_classes))
    model.add(Activation('softmax'))

    model.compile(loss='categorical_crossentropy',
                  optimizer='adam',
                  metrics=['accuracy'])

    # create the model
    # embedding_vecor_length = 32
    # model = Sequential()
    # model.add(Embedding(top_words, embedding_vecor_length, input_length=max_review_length))
    # model.add(Dropout(0.2))
    # model.add(LSTM(100))
    # model.add(Dropout(0.2))
    # model.add(Dense(1, activation='sigmoid'))
    # model.compile(loss='binary_crossentropy', optimizer='adam', metrics=['accuracy'])

    print(model.summary())
    return model


def main():
    data_split = .8
    vocab_size = 5000

    data = load_dataset("train.csv")
    train_input, test_input, train_label, test_label = split_dataset(data, data_split, "Comment", "Insult")
    x_train, x_test, tokenizer = tokenize_input(train_input, test_input, vocab_size)
    y_train, y_test, text_labels, num_classes = encode_labels(train_label, test_label)

    model = create_model(vocab_size, num_classes)
    history = model.fit(x_train, y_train, epochs=10, batch_size=64, verbose=1)

    # Final evaluation of the model
    scores = model.evaluate(x_test, y_test, verbose=1)
    print("Accuracy: %.2f%%" % (scores[1] * 100))

    for i in range(10):
        prediction = model.predict(np.array([x_test[i]]))
        predicted_label = text_labels[np.argmax(prediction)]
        print(test_input.iloc[i][:50], "...")
        print('Actual label:' + str(test_label.iloc[i]))
        print("Predicted label: " + str(predicted_label) + "\n")

    y_softmax = model.predict(x_test)

    y_test_1d = []
    y_pred_1d = []

    for i in range(len(y_test)):
        probs = y_test[i]
        index_arr = np.nonzero(probs)
        one_hot_index = index_arr[0].item(0)
        y_test_1d.append(one_hot_index)

    for i in range(0, len(y_softmax)):
        probs = y_softmax[i]
        predicted_index = np.argmax(probs)
        y_pred_1d.append(predicted_index)

    cnf_matrix = confusion_matrix(y_test_1d, y_pred_1d)
    print(cnf_matrix)
    plt.figure(figsize=(24, 20))
    plot_confusion_matrix(cnf_matrix, classes=text_labels, title="Confusion matrix")
    plt.show()

    while True:
        console_input = input("Enter a message to judge: ")
        input_tokenized = tokenizer.texts_to_matrix([console_input])
        prediction = model.predict(input_tokenized)
        print("Predicted label:", prediction)


if __name__ == '__main__':
    main()
