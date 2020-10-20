import { Bot } from 'barker-api';
import { FC, ReactElement } from 'react';
import { useParams, useHistory } from 'react-router-dom';
import useSWR from 'swr';
import { fetcher, barker } from '../fetcher';

export interface BotLoaderProps {
    render: (props: {
        bot?: Bot;
        error?: any;
        onSubmit: (bot: Bot) => Promise<void>;
    }) => ReactElement;
}
export const BotLoader: FC<BotLoaderProps> = ({ render }) => {
    const { id } = useParams();
    const { data: bot, error, mutate } = useSWR<Bot>(['bot.Get', id], fetcher);
    const onSubmit = async (bot: Bot) => {
        await mutate(barker.bot.Update(bot), false);
    };
    return render({ bot, error, onSubmit });
};
export const NewBotLoader: FC<BotLoaderProps> = ({ render }) => {
    const history = useHistory();
    const onSubmit = async (bot: Bot) => {
        const newBot = await barker.bot.Create(bot);
        history.push(`/bots/${newBot.ID}`);
    };
    return render({ bot: {}, error: null, onSubmit });
};
