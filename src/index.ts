import express from 'express';
import { createServer } from 'http';
import dotenv from 'dotenv';
import router from 'routes/route';
import err_handler from 'middlewares/is_errors';

const app = express();
dotenv.config();


app.use(express.json());
app.use(express.urlencoded({ extended: true }));


const server = createServer(app);

app.use('/api', router);


app.use(err_handler);
server.listen(process.env.PORT, () => {
  console.log(`Server is running on port ${process.env.PORT}`);
});
