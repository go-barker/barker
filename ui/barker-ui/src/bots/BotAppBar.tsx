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
import { Bot } from 'barker-api';
import React, { FC, MouseEvent } from 'react';
import { Link } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        marginBottom: theme.spacing(2),
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
}));

export interface BotAppBarProps {
    bot: Bot;
    isNew?: boolean;
    tab: 'edit' | 'users' | 'campaigns';
}

const BotAppBar: FC<BotAppBarProps> = ({ bot, isNew, tab }) => {
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
                    Bot: {isNew ? '<new>' : bot.Title}
                </Typography>
            </Toolbar>
            <Tabs value={tab}>
                <Tab
                    label="Edit"
                    component={Link}
                    value="edit"
                    to={`/bots/${bot.ID}`}
                />
                {isNew || (
                    <Tab
                        label="Users"
                        component={Link}
                        value="users"
                        to={`/bots/${bot.ID}/users`}
                    />
                )}
                {isNew || (
                    <Tab
                        label="Campaigns"
                        component={Link}
                        value="campaigns"
                        to={`/bots/${bot.ID}/campaigns`}
                    />
                )}
            </Tabs>
        </AppBar>
    );
};

export default BotAppBar;
