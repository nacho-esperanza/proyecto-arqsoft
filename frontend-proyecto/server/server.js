
import express from 'express';

const app = express();

const router = express.Router();

const fetch = (...args) => 
    import('node-fetch').then(({default: fetch}) => fetch(...args));

const itemsBaseURL = "http://localhost:8090/items";

app.get("/api/items/:id", async function(req, res) {

    // codigo hecho en clase

    // Header que permite las llamadas de cualquier dominio para evitar errores.
    res.set('Access-Control-Allow-Origin', '*');

    const url = `${itemsBaseURL}/${req.params.id}`;
    const options = {method: 'GET'};
    try {
        let response = await fetch(url, options);
        let status = response.status;
        response = await response.json();
        res.status(status).json(response);
    } catch (err) {
        console.log(err);
        res.status(500).json({msg: 'Internal Server Error'})
    }
})


app.listen(5001, () => {
    console.log("Server started on port 5001")
});
