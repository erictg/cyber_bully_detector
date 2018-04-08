import itertools

import matplotlib.pyplot as plt
import numpy as np
from keras.layers import Dense, Dropout, Activation
from keras.models import Sequential
from sklearn.metrics import confusion_matrix

# This utility function is from the sklearn docs:
# http://scikit-learn.org/stable/auto_examples/model_selection/plot_confusion_matrix.html
from model_util import tokenize_input, load_dataset, split_dataset, encode_labels, interactive_prompt, init_data


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


def create_model(vocab_size, num_classes):
    model = Sequential()
    model.add(Dense(512, input_shape=(vocab_size,)))
    model.add(Activation('relu'))
    model.add(Dropout(0.5))
    model.add(Dense(num_classes))
    model.add(Activation('softmax'))

    model.compile(loss='categorical_crossentropy',
                  optimizer='adam',
                  metrics=['accuracy'])

    # embedding_vecor_length = 32
    # model = Sequential()
    # model.add(Embedding(vocab_size, embedding_vecor_length))
    # model.add(LSTM(100))
    # model.add(Dropout(.2))
    # model.add(Dense(512, input_shape=(vocab_size,)))
    # model.add(Dense(2, activation='sigmoid'))
    # model.compile(loss='binary_crossentropy', optimizer='adam', metrics=['accuracy'])

    # model = Sequential()
    # model.add(Dense(60, input_shape=(vocab_size,), kernel_initializer='normal', activation='relu'))
    # model.add(Dense(30, kernel_initializer='normal', activation='relu'))
    # model.add(Dropout(.5))
    # model.add(Dense(2, kernel_initializer='normal', activation='sigmoid'))
    # # Compile model
    # model.compile(loss='binary_crossentropy', optimizer='adam', metrics=['accuracy'])

    print(model.summary())
    return model


def main():
    vocab_size = 5000

    train_input, test_input, train_label, test_label, x_train, x_test, y_train, y_test, tokenizer, text_labels, \
    num_classes = init_data("newcsv.csv", vocab_size, .8)

    model = create_model(vocab_size, num_classes)

    model.fit(x_train, y_train, epochs=20, batch_size=64, verbose=1, class_weight={0: 1., 1: 30.})

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

    # model.save('model-' + str(datetime.time()) + '.h5')
    model.save("my_model.h5")  # creates a HDF5 file 'my_model.h5'

    interactive_prompt(model, tokenizer)


if __name__ == '__main__':
    main()
