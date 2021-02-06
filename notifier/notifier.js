const express = require('express');
const notifier = require('node-notifier');
const path = require('path');

const port = process.env.PORT || 8080;

const app = express();

// parsing json-request
app.use(express.json())

app.get('/health', (req, res) => {
    res.status(200).send("notifier healthy");
});

app.post('/notify', (req, res) => {
    notify(req.body, (reply) => {
        res.send(reply);
    });
});

app.listen(port, () => {
    console.log(`server running on port ${port}`);
});

const notify = ({title, message}, cb) => {
    notifier.notify(
        {
            title: title || 'unknown title',
            message: message || 'unknown message',
            icon: path.join(__dirname, 'reminder.png'),
            sound: true,
            wait: true,
            reply: true,
            closeLabel: 'Completed?',
            timeout: 15
        },
        (err, response, reply) => {
            cb(reply);
        },
    );
};
