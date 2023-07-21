//
//  ViewController.m
//  huarongdao-oc
//
//  Created by 銭今墨 on 2023/7/19.
//

#import "ViewController.h"
#include "libhuarongdao.h"

@implementation NavDelegate

- (void)webView:(WKWebView *)webView decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction preferences:(WKWebpagePreferences *)preferences decisionHandler:(void (^)(WKNavigationActionPolicy, WKWebpagePreferences *))decisionHandler{
    decisionHandler(WKNavigationActionPolicyAllow, preferences);
}
- (void)webView:(WKWebView *)webView decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction decisionHandler:(void (^)(WKNavigationActionPolicy))decisionHandler{
    decisionHandler(WKNavigationActionPolicyAllow);
}
- (void)webView:(WKWebView *)webView decidePolicyForNavigationResponse:(WKNavigationResponse *)navigationResponse decisionHandler:(void (^)(WKNavigationResponsePolicy))decisionHandler{
    decisionHandler(WKNavigationResponsePolicyAllow);
}
- (void)webView:(WKWebView *)webView didStartProvisionalNavigation:(WKNavigation *)navigation{}
- (void)webView:(WKWebView *)webView didReceiveServerRedirectForProvisionalNavigation:(WKNavigation *)navigation{}
- (void)webView:(WKWebView *)webView didCommitNavigation:(WKNavigation *)navigation{}
- (void)webView:(WKWebView *)webView didFinishNavigation:(WKNavigation *)navigation{}
- (void)webView:(WKWebView *)webView didReceiveAuthenticationChallenge:(NSURLAuthenticationChallenge *)challenge completionHandler:(void (^)(NSURLSessionAuthChallengeDisposition disposition, NSURLCredential *credential))completionHandler{
    completionHandler(NSURLSessionAuthChallengePerformDefaultHandling, nil);
}
- (void)webView:(WKWebView *)webView authenticationChallenge:(NSURLAuthenticationChallenge *)challenge shouldAllowDeprecatedTLS:(void (^)(BOOL))decisionHandler{
    decisionHandler(YES);
}
- (void)webView:(WKWebView *)webView didFailNavigation:(WKNavigation *)navigation withError:(NSError *)error{}
- (void)webView:(WKWebView *)webView didFailProvisionalNavigation:(WKNavigation *)navigation withError:(NSError *)error{}
- (void)webViewWebContentProcessDidTerminate:(WKWebView *)webView{}
- (void)webView:(WKWebView *)webView navigationResponse:(WKNavigationResponse *)navigationResponse didBecomeDownload:(WKDownload *)download{}
- (void)webView:(WKWebView *)webView navigationAction:(WKNavigationAction *)navigationAction didBecomeDownload:(WKDownload *)download{}

@end

@implementation UIDelegate

- (WKWebView *)webView:(WKWebView *)webView createWebViewWithConfiguration:(WKWebViewConfiguration *)configuration forNavigationAction:(WKNavigationAction *)navigationAction windowFeatures:(WKWindowFeatures *)windowFeatures{
    return nil;
}
- (void)webViewDidClose:(WKWebView *)webView{
    return; nil;
}
@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];

    // Do any additional setup after loading the view.
    WKWebView *webView = [[WKWebView alloc]initWithFrame:self.view.frame];
    NavDelegate *tmpDel =[[NavDelegate alloc]init];
    webView.navigationDelegate=tmpDel;
    UIDelegate *uiDel = [[UIDelegate alloc]init];
    webView.UIDelegate=uiDel;
    [self.view addSubview:webView];
    void* addr = startDarwin();
    NSURL *url =[NSURL URLWithString:[NSString stringWithCString:(const char*)addr encoding:NSASCIIStringEncoding]];
//    NSLog(@"URL=%@", url);
    NSURLRequest *req=[NSURLRequest requestWithURL:url];
//    NSLog(@"REQ=%@", req);
    [webView loadRequest:req];
    freePointer(addr);
}


- (void)setRepresentedObject:(id)representedObject {
    [super setRepresentedObject:representedObject];

    // Update the view, if already loaded.
}


@end
