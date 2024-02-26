import { Request, Response, NextFunction } from 'express';

const err_handler = (err: Error, req: Request, res: Response, next: NextFunction) => {
    const status = res.statusCode !== 200 ? res.statusCode : 500;
    res.json({
        status : status,
        message: err.message,
        stack: process.env.NODE_ENV === 'production' ? '🥞' : err.stack,
    });
}


export default err_handler;
