import { Bot, Campaign } from 'barker-api';
import { FC, ReactElement } from 'react';
import { useParams, useHistory } from 'react-router-dom';
import useSWR from 'swr';
import { fetcher, barker } from '../fetcher';

export interface CampaignLoaderProps {
    render: (props: {
        campaign?: Campaign;
        bot?: Bot;
        error?: any;
        onSubmit: (campaign: Campaign) => Promise<void>;
    }) => ReactElement;
}

export const CampaignLoader: FC<CampaignLoaderProps> = ({ render }) => {
    const { botID, campaignID } = useParams();
    const { data: campaign, error: campaignError, mutate } = useSWR<Campaign>(
        ['campaign.Get', botID, campaignID],
        fetcher
    );
    const { data: bot, error: botError } = useSWR<Bot>(
        ['bot.Get', botID],
        fetcher
    );
    const onSubmit = async (campaign: Campaign) => {
        await mutate(barker.campaign.Update(campaign), false);
    };
    return render({
        campaign,
        bot,
        error: campaignError || botError,
        onSubmit,
    });
};

export const NewCampaignLoader: FC<CampaignLoaderProps> = ({ render }) => {
    const history = useHistory();
    const { botID } = useParams();
    const onSubmit = async (campaign: Campaign) => {
        const newCampaign = await barker.campaign.Create(campaign);
        history.push(`/bots/${newCampaign.BotID}/campaigns/${newCampaign.ID}`);
    };
    const { data: bot, error: botError } = useSWR<Bot>(
        ['bot.Get', botID],
        fetcher
    );
    return render({
        campaign: { BotID: parseInt(botID, 10) },
        bot,
        error: botError,
        onSubmit,
    });
};
