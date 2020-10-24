import { Grid, makeStyles, Paper } from '@material-ui/core';
import { Campaign, Bot } from 'barker-api';
import React, { FC } from 'react';
import BotAppBar from '../bots/BotAppBar';
import CampaignEditForm from './CampaignEditForm';

const useStyles = makeStyles((theme) => ({
    paper: {
        padding: theme.spacing(2),
    },
}));

export interface CampaignPageProps {
    campaign?: Campaign;
    bot?: Bot;
    error?: any;
    isNew?: boolean;
    onSubmit: (campaign: Campaign) => Promise<void>;
}
export const CampaignPage: FC<CampaignPageProps> = ({
    campaign,
    error,
    isNew,
    onSubmit,
    bot,
}) => {
    const classes = useStyles();

    if (error) {
        return <div>failed to load</div>;
    }
    if (!campaign || !bot) return <div>loading...</div>;

    return (
        <Grid container>
            <BotAppBar
                botID={bot.ID ?? 0}
                tab="campaigns"
                title={
                    isNew
                        ? 'New campaign'
                        : `Campaign: ${campaign.Title || '<untitled>'}`
                }
            />

            <Paper className={classes.paper}>
                <CampaignEditForm onSubmit={onSubmit} campaign={campaign} />
            </Paper>
        </Grid>
    );
};

export default CampaignPage;
