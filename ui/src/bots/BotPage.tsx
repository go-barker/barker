import { Grid, makeStyles, Paper } from '@material-ui/core';
import { Bot } from 'barker-api';
import React, { FC } from 'react';
import { UsersListLoader } from '../users/UsersListLoader';
import { UsersListPage } from '../users/UsersListPage';
import BotAppBar from './BotAppBar';
import BotEditForm from './BotEditForm';
import { CampaignsListLoader } from '../campaigns/CampaignsListLoader';
import { CampaignsListPage } from '../campaigns/CampaignsListPage';

const useStyles = makeStyles((theme) => ({
    paper: {
        padding: theme.spacing(2),
        width: '100%',
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
            <BotAppBar
                botID={bot.ID ?? 0}
                isNew={isNew}
                tab={tab}
                title={isNew ? 'New bot' : `Bot: ${bot.Title || '<untitled>'}`}
            />

            <Paper className={classes.paper}>
                {tab === 'edit' && (
                    <BotEditForm onSubmit={onSubmit} bot={bot} />
                )}
                {tab === 'users' && (
                    <UsersListLoader
                        render={(props) => <UsersListPage {...props} />}
                    />
                )}
                {tab === 'campaigns' && (
                    <CampaignsListLoader
                        render={(props) => <CampaignsListPage {...props} />}
                    />
                )}
            </Paper>
        </Grid>
    );
};

export default BotPage;
