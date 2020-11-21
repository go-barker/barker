import { Grid, makeStyles, Paper, Typography } from '@material-ui/core';
import { Bot, Campaign } from 'barker-api';
import React, { FC } from 'react';
import BotAppBar from '../bots/BotAppBar';
import CampaignEditForm from './CampaignEditForm';
import { CampaignStatWidget } from './CampaignStatWidget';

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
                <Grid container spacing={2}>
                    <Grid item xs={12} sm={6}>
                        <CampaignEditForm
                            onSubmit={onSubmit}
                            campaign={campaign}
                        />
                    </Grid>
                    {!isNew && campaign.BotID && campaign.ID && (
                        <Grid item xs={12} sm={6}>
                            <Typography variant="subtitle1">
                                Progress
                            </Typography>
                            <CampaignStatWidget
                                botID={campaign.BotID}
                                campaignID={campaign.ID}
                            />
                        </Grid>
                    )}
                </Grid>
            </Paper>
        </Grid>
    );
};

export default CampaignPage;
