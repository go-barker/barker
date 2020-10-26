import { Grid, makeStyles } from '@material-ui/core';
import React, { FC } from 'react';
import texture from './grey.png';

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundImage: `url(${texture})`,
        height: '100vh',
    },
    wrapper: {
        width: theme.breakpoints.values.md,
        [theme.breakpoints.down('sm')]: {
            width: '100%',
        },
    },
}));

export interface LayoutProps {}
export const Layout: FC<LayoutProps> = ({ children }) => {
    const classes = useStyles();
    return (
        <Grid container justify="center" className={classes.root}>
            <Grid item className={classes.wrapper}>
                {children}
            </Grid>
        </Grid>
    );
};
