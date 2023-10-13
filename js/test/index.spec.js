import { expect } from "chai";
import { gameOfLifeJs } from "../app.js";

describe('gameOfLifeJs', () => {
  it('should return Hello Dojoers !', () => {
    expect(gameOfLifeJs()).to.be.equal('Hello Dojoers !')
  });
});