import { Angular2GinPage } from './app.po';

describe('angular2-gin App', function() {
  let page: Angular2GinPage;

  beforeEach(() => {
    page = new Angular2GinPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
