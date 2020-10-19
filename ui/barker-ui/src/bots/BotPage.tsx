import { Button, Grid, Typography, Paper, makeStyles } from '@material-ui/core';
import { Bot } from 'barker-api';
import { Field, Form, Formik } from 'formik';
import { TextField } from 'formik-material-ui';
import React, { FC } from 'react';
import { mutate } from 'swr';
import { barker } from '../fetcher';
import { useHistory } from 'react-router-dom';
import BotAppBar from './BotAppBar';
import BotEditForm from './BotEditForm';

const useStyles = makeStyles((theme) => ({
    paper: {
        padding: theme.spacing(2),
    },
}));

export interface BotPageProps {
    bot?: Bot;
    tab: 'edit' | 'users' | 'campaigns';
    error?: any;
    isNew?: boolean;
    onSubmit: (bot: Bot) => Promise<void>;
}
export const BotPage: FC<BotPageProps> = ({
    bot,
    tab,
    error,
    isNew,
    onSubmit,
}) => {
    const classes = useStyles();

    if (error) {
        return <div>failed to load</div>;
    }
    if (!bot) return <div>loading...</div>;

    return (
        <Grid container>
            <BotAppBar bot={bot} isNew={isNew} tab={tab} />

            <Paper className={classes.paper}>
                {tab === 'edit' && (
                    <BotEditForm onSubmit={onSubmit} bot={bot} />
                )}
            </Paper>
        </Grid>
    );
};

export default BotPage;
