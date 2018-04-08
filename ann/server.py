from flask import Flask, request, jsonify
import json
import cleaner
import model_util

app = Flask(__name__)


@app.route("/classify", methods=['POST'])
def classify():
    data = request.get_json()
    input_tokenized = tokenizer.texts_to_matrix([cleaner.clean_message(data['content'])])
    prediction = model.predict(input_tokenized)[0]

    prediction_dict = {'not_insult': float(prediction[0]), 'insult': float(prediction[1])}

    return jsonify(prediction_dict)


if __name__ == '__main__':
    _, _, _, _, _, _, _, _, tokenizer, _, _ = model_util.init_data("newcsv.csv", 5000, .8)
    model = model_util.load_modelfile("single_big_dense_layer.h5")
    app.run()
