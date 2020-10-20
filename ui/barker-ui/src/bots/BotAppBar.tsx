import { Bot } from 'barker-api';
import React, { FC } from 'react';
import NavigationBar, { NavigationBarTab } from '../NavigationBar';

export interface BotAppBarProps {
    bot: Bot;
    isNew?: boolean;
    tab: 'edit' | 'users' | 'campaigns';
}

const BotAppBar: FC<BotAppBarProps> = ({ bot, isNew, tab }) => {
    const tabs: NavigationBarTab[] = [
        {
            label: 'Bots',
            href: `/`,
            value: 'bots',
        },
        {
            label: 'Edit',
            href: `/bots/${bot.ID}`,
            value: 'edit',
        },
    ];
    if (!isNew) {
        tabs.push(
            {
                href: `/bots/${bot.ID}/users`,
                label: 'Users',
                value: 'users',
            },
            {
                href: `/bots/${bot.ID}/campaigns`,
                label: 'Campaigns',
                value: 'campaigns',
            }
        );
    }
    return (
        <NavigationBar
            tabs={tabs}
            tab={tab}
            title={'Bot: ' + (isNew ? '<new>' : bot.Title || '<untitled>')}
        />
    );
};

export default BotAppBar;
