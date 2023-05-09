const { Kafka } = require('kafkajs')
const http = require('http');
const { getReqData } = require('./utils');
const crypto = require('crypto');

const PORT = process.env.PORT || 3003;
const TOPIC = 'eshop';
const cbs = {};

const kafka = new Kafka({
  clientId: 'app-proxy',
  brokers: ['kafka:9092'],
})

const producer = kafka.producer();
const consumer = kafka.consumer({ groupId: 'app-proxy-consumer' });

process.on('beforeExit', async () => {
  console.log('Disconnecting kafka producer');
  await producer.disconnect();
});


(async () => {
  await producer.connect();
  await consumer.connect();
  await consumer.subscribe({ topic: TOPIC });

  await consumer.run({
    eachMessage: async ({ topic, message }) => {
      if (topic !== TOPIC)
        return;

      const [ms, id, data] = message.value.toString().split(':@?@:');
      if (ms !== 'viewer')
        return;

      const cb = cbs[id];
      if (cb) {
        cb(data);
        delete cbs[id];
      }
    },
  });


  const server = http.createServer(async (req, res) => {
    if (req.url.match(/\/proxy\/([a-z-]+)/) && req.method === 'POST') {
      const ms = req.url.split('/')[2];
      const data = await getReqData(req);
      const id = crypto.randomBytes(32).toString('hex');

      const to = setTimeout(() => {
        res.writeHead(408, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ message: 'Request Timeout' }));
        delete cbs[id];
      }, 30 * 1000); // 30 seconds

      cbs[id] = (data) => {
        clearTimeout(to);
        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(data);
      }

      await producer.send({
        topic: TOPIC,
        messages: [{ value: [ms, id, data].join(':@?@:') }],
      });

    } else {
      res.writeHead(404, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ message: 'Route not found' }));
    }
  });

  server.listen(PORT, () => {
    console.log(`server started on port: ${PORT}`);
  });


})();
