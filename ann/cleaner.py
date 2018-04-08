from nltk.tokenize import word_tokenize
import csv


def clean_message(text):
    text = text.replace('_', ' ')
    text = text.replace('-', ' ')
    tokens = word_tokenize(text)
    words = [word for word in tokens if word.isalpha()]
    new_words = []
    for word in words:
        new_words.append(word.lower())

    return ' '.join(new_words)


def clean_csv(input_filename, output_filename):
    csvfile = open(input_filename, 'r')
    reader = csv.reader(csvfile, quotechar='|')
    output_csvfile = open(output_filename, 'w')
    writer = csv.writer(output_csvfile, quotechar='|')

    first_row = True

    for row in reader:
        if first_row:
            writer.writerow(row)
            first_row = False
        else:
            writer.writerow([row[0], row[1], clean_message(row[2])])

    csvfile.close()
    output_csvfile.close()

if __name__ == '__main__':
    clean_csv("train.csv", "newcsv.csv")