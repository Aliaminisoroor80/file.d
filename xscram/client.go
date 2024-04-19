@startuml
from flask import Flask, render_template, request, redirect, url_for
from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, SubmitField
from wtforms.validators import InputRequired, Email, Length
from flask_wtf.file import FileField, FileRequired
from flask_pdf import PDF

app = Flask(__name__)
app.config['SECRET_KEY'] = 'your_secret_key'

pdf = PDF(app)

class LoginForm(FlaskForm):
    email = StringField('email', validators=[InputRequired(), Email(message='The email is not valid')])
    password = PasswordField('password', validators=[InputRequired(), Length(min=6, message='Password must be at least 6 characters long')])
    submit = SubmitField('log in')

class RegistrationForm(FlaskForm):
    email = StringField('email', validators=[InputRequired(), Email(message='The email is not valid')])
    password = PasswordField('password', validators=[InputRequired(), Length(min=6, message='Password must be at least 6 characters long')])
    photo = FileField('Upload profile picture', validators=[FileRequired()])
    submit = SubmitField('Register')

@app.route('/login', methods=['GET', 'POST'])
def login():
    form = LoginForm()
    if form.validate_on_submit():
        # Validation of user login information
        # This part has been omitted for simplicity
        return redirect(url_for('index'))
    return render_template('login.html', form=form)

@app.route('/register', methods=['GET', 'POST'])
def register():
    form = RegistrationForm()
    if form.validate_on_submit():
        # Store user information in the database
        # This part has been omitted for simplicity
        return redirect(url_for('index'))
    return render_template('register.html', form=form)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/generate_pdf', methods=['GET', 'POST'])
def generate_pdf():
    data = {'username': 'John Doe', 'email': 'john@example.com'}  # Data to be displayed in PDF
    rendered = render_template('pdf_template.html', data=data)
    pdf.from_string(rendered, 'pdf_output.pdf')
    return 'PDF created successfully. <a href="/static/pdf_output.pdf">Download PDF</a>'

if __name__ == '__main__':
    app.run(debug=True)
```
