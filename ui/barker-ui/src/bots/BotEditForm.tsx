import { Button, Grid } from '@material-ui/core';
import { Bot } from 'barker-api';
import { Field, Form, Formik } from 'formik';
import { TextField } from 'formik-material-ui';
import React, { FC } from 'react';

export interface BotEditFormProps {
    bot: Bot;
    onSubmit: (bot: Bot) => Promise<void>;
}

const BotEditForm: FC<BotEditFormProps> = ({ bot, onSubmit }) => {
    return (
        <Grid container>
            <Grid item>
                <Formik initialValues={bot} onSubmit={onSubmit}>
                    <Form>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <Field
                                    disabled
                                    component={TextField}
                                    label="ID"
                                    name="ID"
                                    variant="outlined"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Field
                                    component={TextField}
                                    label="Title"
                                    name="Title"
                                    variant="outlined"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Field
                                    component={TextField}
                                    label="Token"
                                    name="Token"
                                    variant="outlined"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Button
                                    type="submit"
                                    color="primary"
                                    variant="contained"
                                >
                                    Save
                                </Button>
                            </Grid>
                        </Grid>
                    </Form>
                </Formik>
            </Grid>
        </Grid>
    );
};
export default BotEditForm;
