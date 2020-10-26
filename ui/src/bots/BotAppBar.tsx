import React, { FC } from 'react';
import NavigationBar, { NavigationBarTab } from '../NavigationBar';

export interface BotAppBarProps {
    botID: number;
    isNew?: boolean;
    title: string;
    tab: 'edit' | 'users' | 'campaigns';
}

const BotAppBar: FC<BotAppBarProps> = ({ botID, isNew, tab, title }) => {
    const tabs: NavigationBarTab[] = [
        {
            label: 'Bots',
            href: `/`,
            value: 'bots',
        },
        {
            label: 'Bot settings',
            href: `/bots/${botID}`,
            value: 'edit',
        },
    ];
    if (!isNew) {
        tabs.push(
            {
                href: `/bots/${botID}/users`,
                label: 'Users',
                value: 'users',
            },
            {
                href: `/bots/${botID}/campaigns`,
                label: 'Campaigns',
                value: 'campaigns',
            }
        );
    }
    return <NavigationBar tabs={tabs} tab={tab} title={title} />;
};

export default BotAppBar;
