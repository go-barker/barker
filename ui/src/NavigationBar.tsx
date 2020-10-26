import {
    AppBar,
    IconButton,
    makeStyles,
    Tab,
    Tabs,
    Toolbar,
    Typography,
} from '@material-ui/core';
import { Menu as MenuIcon } from '@material-ui/icons';
import React, { FC } from 'react';
import { Link } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
}));
export interface NavigationBarTab {
    href: string;
    label: string;
    value: string;
}
export interface NavigationBarProps {
    title: string;
    tab: string;
    tabs: NavigationBarTab[];
}

const NavigationBar: FC<NavigationBarProps> = ({ tab, tabs, title }) => {
    const classes = useStyles();
    return (
        <AppBar position="static" className={classes.root}>
            <Toolbar>
                <IconButton
                    edge="start"
                    className={classes.menuButton}
                    color="inherit"
                    aria-label="menu"
                >
                    <MenuIcon />
                </IconButton>
                <Typography variant="h6" className={classes.title}>
                    {title}
                </Typography>
            </Toolbar>
            <Tabs value={tab}>
                {tabs.map((tab, i) => (
                    <Tab
                        key={i}
                        label={tab.label}
                        component={Link}
                        value={tab.value}
                        to={tab.href}
                    />
                ))}
            </Tabs>
        </AppBar>
    );
};

export default NavigationBar;
