from flask import Flask


from .config import Config
from .models import db
from .book_routes import main_routes

def create_app():
    app = Flask(__name__)
    app.config.from_object(Config)  
    db.init_app(app)   

    
    app.register_blueprint(main_routes) 
    return app