export interface MarketTick {
  symbol: string;
  price: number;
  timestamp: number;
  volume: number;
}

export interface TradePosition {
  id: string;
  asset: string;
  type: 'CALL' | 'PUT';
  entryPrice: number;
  investment: number;
  payoutPercent: number;
  expiryTimestamp: number;
}

export interface TradingEngineContract {
  subscribeTicks(symbol: string, callback: (tick: MarketTick) => void): () => void;
  executeOrder(order: Omit<TradePosition, 'id'>): Promise<{ success: boolean; positionId: string }>;
}
