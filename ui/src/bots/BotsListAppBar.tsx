import React, { FC } from 'react';
import NavigationBar, { NavigationBarTab } from '../NavigationBar';

export interface BotsListAppBarProps {}

const BotsListAppBar: FC<BotsListAppBarProps> = () => {
    const tabs: NavigationBarTab[] = [
        {
            label: 'Bots',
            href: `/`,
            value: 'bots',
        },
        {
            label: 'Create bot',
            href: `/bots/new`,
            value: 'create',
        },
    ];
    return <NavigationBar tabs={tabs} tab="bots" title="Bots" />;
};

export default BotsListAppBar;
