#!/usr/bin/env python

from flask import Flask, request, render_template, make_response


app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'


@app.route('/pac')
def pac_main():
    socket_host_default = '127.0.0.1'
    socket_port_default = 1080
    socket_host = request.args.get('host', socket_host_default)
    socket_port = request.args.get('port', socket_port_default)
    text = render_template('pac.template', host=socket_host, port=socket_port)
    resp = make_response(text)
    resp.headers['Content-Type'] = 'text/plain'
    return resp

if __name__ == '__main__':
    app.run(host='0.0.0.0')


