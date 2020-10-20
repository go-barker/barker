import { Bot } from 'barker-api';
import { createListLoader } from '../createListLoader';

export const BotsListLoader = createListLoader<Bot>({ key: 'bot.List' });
