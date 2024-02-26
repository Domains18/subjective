import express from 'express';
import { createServer } from 'http';
import dotenv from 'dotenv';
import router from './routes/router';
import {err_handler} from './middlewares/is_errors';


import {port }from './constants/variables';

const app = express();
dotenv.config();


app.use(express.json());
app.use(express.urlencoded({ extended: true }));


const server = createServer(app);

app.use('/api', router());


app.use(err_handler);
server.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
