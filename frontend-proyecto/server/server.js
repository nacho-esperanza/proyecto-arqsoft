
import express from 'express';

const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const app = express();


// Middleware
app.use(bodyParser.json());
app.use(cors());

const router = express.Router();

// Rutas

//Ruta get para obtener el usuario por ID
router.get('/users/:id', (req, res) => {
    const id = req.params.id;
    const user = users.find(user => user.id == id);
    res.send(user);
});

//Ruta get para obtener todos los usuarios
router.get('/users', (req, res) => {
    res.send(users);
});

//Ruta post para crear un usuario
router.post('/users', (req, res) => {
    const user = req.body;
    user.id = users.length + 1;
    users.push(user);
    res.send(user);
});

//Ruta post para loguear un usuario
router.post('/user/login', (req, res) => {
    const user = req.body;
    const userFound = users.find(u => u.email == user.email && u.password == user.password);
    if (userFound) {
        res.send(userFound);
    } else {
        res.status(401).send({msg: 'Usuario o contraseña incorrectos'});
    }
});
